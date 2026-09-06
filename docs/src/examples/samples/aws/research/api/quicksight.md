# Amazon QuickSight

API version: 2018-04-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/quicksight/2018-04-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateTopicReviewedAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `Answers` | `List<CreateTopicReviewedAnswer>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `SucceededAnswers` | `List<SucceededTopicReviewedAnswer>` | no |
| `InvalidAnswers` | `List<InvalidTopicReviewedAnswer>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## BatchDeleteKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Deleted` | `List<BatchDeleteKnowledgeBaseSuccess>` | yes |
| `Errors` | `List<BatchDeleteKnowledgeBaseFailure>` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## BatchDeleteTopicReviewedAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `AnswerIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `SucceededAnswers` | `List<SucceededTopicReviewedAnswer>` | no |
| `InvalidAnswers` | `List<InvalidTopicReviewedAnswer>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## BatchDescribeUserLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `users` | `List<UserLimitsEntry>` | no |
| `resourceTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userLimits` | `List<UserLimits>` | yes |
| `errors` | `List<BatchDescribeUserLimitsError>` | yes |

## CancelIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `IngestionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `IngestionId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateAccountCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |
| `AccountCustomization` | `AccountCustomization` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AwsAccountId` | `string` | no |
| `Namespace` | `string` | no |
| `AccountCustomization` | `AccountCustomization` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateAccountSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Edition` | `string` | no |
| `AuthenticationMethod` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `AccountName` | `string` | yes |
| `NotificationEmail` | `string` | yes |
| `ActiveDirectoryName` | `string` | no |
| `Realm` | `string` | no |
| `DirectoryId` | `string` | no |
| `AdminGroup` | `List<string>` | no |
| `AuthorGroup` | `List<string>` | no |
| `ReaderGroup` | `List<string>` | no |
| `AdminProGroup` | `List<string>` | no |
| `AuthorProGroup` | `List<string>` | no |
| `ReaderProGroup` | `List<string>` | no |
| `FirstName` | `string` | no |
| `LastName` | `string` | no |
| `EmailAddress` | `string` | no |
| `ContactNumber` | `string` | no |
| `IAMIdentityCenterInstanceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SignupResponse` | `SignupResponse` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateActionConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ActionConnectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `AuthenticationConfig` | `AuthConfig` | yes |
| `Description` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `VpcConnectionArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationStatus` | `string` | no |
| `ActionConnectorId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Spaces` | `List<string>` | no |
| `ActionConnectors` | `List<string>` | no |
| `AwsAccountId` | `string` | yes |
| `AgentId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `IconId` | `string` | no |
| `StarterPrompts` | `List<string>` | no |
| `WelcomeMessage` | `string` | no |
| `AgentLifecycle` | `string` | no |
| `CustomPromptInput` | `CustomPromptInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `AgentId` | `string` | yes |
| `AgentStatus` | `string` | yes |
| `AgentName` | `string` | yes |
| `RequestId` | `string` | no |

## CreateAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |
| `Name` | `string` | yes |
| `Parameters` | `Parameters` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `SourceEntity` | `AnalysisSourceEntity` | no |
| `ThemeArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Definition` | `AnalysisDefinition` | no |
| `ValidationStrategy` | `ValidationStrategy` | no |
| `FolderArns` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AnalysisId` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateApprovalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Actions` | `List<string>` | yes |
| `AssetTypes` | `List<string>` | yes |
| `ApplicableTo` | `ApplicableTo` | yes |
| `ApprovalGroups` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `ApprovalPolicy` | yes |

## CreateBrand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandId` | `string` | yes |
| `BrandDefinition` | `BrandDefinition` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `BrandDetail` | `BrandDetail` | no |
| `BrandDefinition` | `BrandDefinition` | no |

## CreateCustomPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `CustomPermissionsName` | `string` | yes |
| `Capabilities` | `Capabilities` | no |
| `Governance` | `Governance` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `RequestId` | `string` | no |

## CreateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `Name` | `string` | yes |
| `Parameters` | `Parameters` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `SourceEntity` | `DashboardSourceEntity` | no |
| `Tags` | `List<Tag>` | no |
| `VersionDescription` | `string` | no |
| `DashboardPublishOptions` | `DashboardPublishOptions` | no |
| `ThemeArn` | `string` | no |
| `Definition` | `DashboardVersionDefinition` | no |
| `ValidationStrategy` | `ValidationStrategy` | no |
| `FolderArns` | `List<string>` | no |
| `LinkSharingConfiguration` | `LinkSharingConfiguration` | no |
| `LinkEntities` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VersionArn` | `string` | no |
| `DashboardId` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `Name` | `string` | yes |
| `PhysicalTableMap` | `Map<PhysicalTable>` | yes |
| `LogicalTableMap` | `Map<LogicalTable>` | no |
| `ImportMode` | `string` | yes |
| `ColumnGroups` | `List<ColumnGroup>` | no |
| `FieldFolders` | `Map<FieldFolder>` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RowLevelPermissionDataSet` | `RowLevelPermissionDataSet` | no |
| `RowLevelPermissionTagConfiguration` | `RowLevelPermissionTagConfiguration` | no |
| `ColumnLevelPermissionRules` | `List<ColumnLevelPermissionRule>` | no |
| `Tags` | `List<Tag>` | no |
| `DataSetUsageConfiguration` | `DataSetUsageConfiguration` | no |
| `DatasetParameters` | `List<DatasetParameter>` | no |
| `FolderArns` | `List<string>` | no |
| `PerformanceConfiguration` | `PerformanceConfiguration` | no |
| `UseAs` | `string` | no |
| `DataPrepConfiguration` | `DataPrepConfiguration` | no |
| `SemanticModelConfiguration` | `SemanticModelConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `DataSetId` | `string` | no |
| `IngestionArn` | `string` | no |
| `IngestionId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSourceId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `DataSourceParameters` | `DataSourceParameters` | no |
| `Credentials` | `DataSourceCredentials` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `VpcConnectionProperties` | `VpcConnectionProperties` | no |
| `SslProperties` | `SslProperties` | no |
| `Tags` | `List<Tag>` | no |
| `FolderArns` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `DataSourceId` | `string` | no |
| `CreationStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateDlpSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DlpSettingId` | `string` | yes |
| `Name` | `string` | yes |
| `ProviderType` | `string` | yes |
| `ProviderConfig` | `ProviderConfig` | yes |
| `ProviderOutageAction` | `string` | yes |
| `Enabled` | `boolean` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `DlpSettingId` | `string` | yes |
| `RequestId` | `string` | no |

## CreateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `FlowDefinition` | `SensitiveDocument` | yes |
| `Permissions` | `List<Permission>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `FlowId` | `string` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `Name` | `string` | no |
| `FolderType` | `string` | no |
| `ParentFolderArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Tags` | `List<Tag>` | no |
| `SharingModel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `FolderId` | `string` | no |
| `RequestId` | `string` | no |

## CreateFolderMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `MemberId` | `string` | yes |
| `MemberType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `FolderMember` | `FolderMember` | no |
| `RequestId` | `string` | no |

## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `Description` | `string` | no |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateGroupMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberName` | `string` | yes |
| `GroupName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupMember` | `GroupMember` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateIAMPolicyAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssignmentName` | `string` | yes |
| `AssignmentStatus` | `string` | yes |
| `PolicyArn` | `string` | no |
| `Identities` | `Map<List<string>>` | no |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentName` | `string` | no |
| `AssignmentId` | `string` | no |
| `AssignmentStatus` | `string` | no |
| `PolicyArn` | `string` | no |
| `Identities` | `Map<List<string>>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `IngestionId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `IngestionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `IngestionId` | `string` | no |
| `IngestionStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `Name` | `string` | yes |
| `DataSourceArn` | `string` | yes |
| `KnowledgeBaseConfiguration` | `KnowledgeBaseConfiguration` | yes |
| `Description` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `MediaExtractionConfiguration` | `MediaExtractionConfiguration` | no |
| `AccessControlConfiguration` | `AccessControlConfiguration` | no |
| `PrimaryOwnerArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseArn` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `CreationStatus` | `string` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateLimitsProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `profileName` | `string` | yes |
| `description` | `string` | no |
| `resourceLimits` | `Map<ProfileLimitValue>` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `profileId` | `string` | yes |

## CreateNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `IdentityStore` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `CapacityRegion` | `string` | no |
| `CreationStatus` | `string` | no |
| `IdentityStore` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateOAuthClientApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `OAuthClientApplicationId` | `string` | yes |
| `Name` | `string` | yes |
| `OAuthClientAuthenticationType` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientSecret` | `string` | yes |
| `OAuthTokenEndpointUrl` | `string` | yes |
| `OAuthAuthorizationEndpointUrl` | `string` | no |
| `OAuthScopes` | `string` | no |
| `DataSourceType` | `string` | no |
| `IdentityProviderVpcConnectionProperties` | `VpcConnectionProperties` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `OAuthClientApplicationId` | `string` | no |
| `CreationStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Schedule` | `RefreshSchedule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `ScheduleId` | `string` | no |
| `Arn` | `string` | no |

## CreateRoleMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `Role` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `RequestId` | `string` | no |

## CreateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `Name` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `SourceEntity` | `TemplateSourceEntity` | no |
| `Tags` | `List<Tag>` | no |
| `VersionDescription` | `string` | no |
| `Definition` | `TemplateVersionDefinition` | no |
| `ValidationStrategy` | `ValidationStrategy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VersionArn` | `string` | no |
| `TemplateId` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateTemplateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `AliasName` | `string` | yes |
| `TemplateVersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateAlias` | `TemplateAlias` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `Name` | `string` | yes |
| `BaseThemeId` | `string` | yes |
| `VersionDescription` | `string` | no |
| `Configuration` | `ThemeConfiguration` | yes |
| `Permissions` | `List<ResourcePermission>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VersionArn` | `string` | no |
| `ThemeId` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateThemeAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `AliasName` | `string` | yes |
| `ThemeVersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeAlias` | `ThemeAlias` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `Topic` | `TopicDetails` | yes |
| `Tags` | `List<Tag>` | no |
| `FolderArns` | `List<string>` | no |
| `CustomInstructions` | `CustomInstructions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `RefreshArn` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateTopicRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `DatasetArn` | `string` | yes |
| `DatasetName` | `string` | no |
| `RefreshSchedule` | `TopicRefreshSchedule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `DatasetArn` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## CreateTopicV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `Topic` | `TopicV2Details` | yes |
| `Tags` | `List<Tag>` | no |
| `FolderArns` | `List<string>` | no |
| `CustomInstructions` | `CustomInstructions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## CreateVPCConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `VPCConnectionId` | `string` | yes |
| `Name` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `SecurityGroupIds` | `List<string>` | yes |
| `DnsResolvers` | `List<string>` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VPCConnectionId` | `string` | no |
| `CreationStatus` | `string` | no |
| `AvailabilityStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteAccountCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteAccountCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteAccountSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteActionConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ActionConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ActionConnectorId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentId` | `string` | yes |
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |

## DeleteAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |
| `RecoveryWindowInDays` | `long` | no |
| `ForceDeleteWithoutRecovery` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `AnalysisId` | `string` | no |
| `DeletionTime` | `timestamp` | no |
| `RequestId` | `string` | no |

## DeleteApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |

## DeleteApprovalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBrand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |

## DeleteBrandAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |

## DeleteCustomPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `CustomPermissionsName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `RequestId` | `string` | no |

## DeleteDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `VersionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `DashboardId` | `string` | no |
| `RequestId` | `string` | no |

## DeleteDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `DataSetId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteDataSetRefreshProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `DataSourceId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteDefaultQBusinessApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteDlpSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DlpSettingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `DlpSettingId` | `string` | yes |
| `RequestId` | `string` | no |

## DeleteFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `FolderId` | `string` | no |
| `RequestId` | `string` | no |

## DeleteFolderMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `MemberId` | `string` | yes |
| `MemberType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteGroupMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberName` | `string` | yes |
| `GroupName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteIAMPolicyAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssignmentName` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentName` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteIdentityPropagationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Service` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseArn` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteLimitsProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## DeleteNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteOAuthClientApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `OAuthClientApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `OAuthClientApplicationId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `ScheduleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `ScheduleId` | `string` | no |
| `Arn` | `string` | no |

## DeleteRoleCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteRoleMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberName` | `string` | yes |
| `Role` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `RequestId` | `string` | no |

## DeleteTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `VersionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Arn` | `string` | no |
| `TemplateId` | `string` | no |
| `Status` | `integer` | no |

## DeleteTemplateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `TemplateId` | `string` | no |
| `AliasName` | `string` | no |
| `Arn` | `string` | no |
| `RequestId` | `string` | no |

## DeleteTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `VersionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `ThemeId` | `string` | no |

## DeleteThemeAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | no |
| `Arn` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `ThemeId` | `string` | no |

## DeleteTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteTopicRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `DatasetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `DatasetArn` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DeleteTopicV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteUserByPrincipalId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrincipalId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteUserCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DeleteVPCConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `VPCConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VPCConnectionId` | `string` | no |
| `DeletionStatus` | `string` | no |
| `AvailabilityStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeAccountCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomPermissionsName` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeAccountCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |
| `Resolved` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AwsAccountId` | `string` | no |
| `Namespace` | `string` | no |
| `AccountCustomization` | `AccountCustomization` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountSettings` | `AccountSettings` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeAccountSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountInfo` | `AccountInfo` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeActionConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ActionConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionConnector` | `ActionConnector` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeActionConnectorPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ActionConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ActionConnectorId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentId` | `string` | yes |
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Agent` | `Agent` | yes |
| `RequestId` | `string` | no |

## DescribeAgentPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentId` | `string` | yes |
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `AgentId` | `string` | yes |
| `Permissions` | `List<ResourcePermission>` | yes |
| `RequestId` | `string` | yes |

## DescribeAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Analysis` | `Analysis` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeAnalysisDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisId` | `string` | no |
| `Name` | `string` | no |
| `Errors` | `List<AnalysisError>` | no |
| `ResourceStatus` | `string` | no |
| `ThemeArn` | `string` | no |
| `Definition` | `AnalysisDefinition` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeAnalysisPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisId` | `string` | no |
| `AnalysisArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `App` | `AppSummary` | yes |
| `RequestId` | `string` | no |

## DescribeAppPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `Arn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |

## DescribeApprovalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `ApprovalPolicy` | yes |

## DescribeAssetBundleExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssetBundleExportJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `DownloadUrl` | `string` | no |
| `Errors` | `List<AssetBundleExportJobError>` | no |
| `Arn` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `AssetBundleExportJobId` | `string` | no |
| `AwsAccountId` | `string` | no |
| `ResourceArns` | `List<string>` | no |
| `IncludeAllDependencies` | `boolean` | no |
| `ExportFormat` | `string` | no |
| `CloudFormationOverridePropertyConfiguration` | `AssetBundleCloudFormationOverridePropertyConfiguration` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `IncludePermissions` | `boolean` | no |
| `IncludeTags` | `boolean` | no |
| `ValidationStrategy` | `AssetBundleExportJobValidationStrategy` | no |
| `Warnings` | `List<AssetBundleExportJobWarning>` | no |
| `IncludeFolderMemberships` | `boolean` | no |
| `IncludeFolderMembers` | `string` | no |

## DescribeAssetBundleImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssetBundleImportJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `Errors` | `List<AssetBundleImportJobError>` | no |
| `RollbackErrors` | `List<AssetBundleImportJobError>` | no |
| `Arn` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `AssetBundleImportJobId` | `string` | no |
| `AwsAccountId` | `string` | no |
| `AssetBundleImportSource` | `AssetBundleImportSourceDescription` | no |
| `OverrideParameters` | `AssetBundleImportJobOverrideParameters` | no |
| `FailureAction` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `OverridePermissions` | `AssetBundleImportJobOverridePermissions` | no |
| `OverrideTags` | `AssetBundleImportJobOverrideTags` | no |
| `OverrideValidationStrategy` | `AssetBundleImportJobOverrideValidationStrategy` | no |
| `Warnings` | `List<AssetBundleImportJobWarning>` | no |

## DescribeAutomationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AutomationGroupId` | `string` | yes |
| `AutomationId` | `string` | yes |
| `IncludeInputPayload` | `boolean` | no |
| `IncludeOutputPayload` | `boolean` | no |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | no |
| `StartedAt` | `timestamp` | no |
| `EndedAt` | `timestamp` | no |
| `JobStatus` | `string` | yes |
| `InputPayload` | `string` | no |
| `OutputPayload` | `string` | no |
| `RequestId` | `string` | no |

## DescribeBrand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandId` | `string` | yes |
| `VersionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `BrandDetail` | `BrandDetail` | no |
| `BrandDefinition` | `BrandDefinition` | no |

## DescribeBrandAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `BrandArn` | `string` | no |

## DescribeBrandPublishedVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `BrandDetail` | `BrandDetail` | no |
| `BrandDefinition` | `BrandDefinition` | no |

## DescribeCustomPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `CustomPermissionsName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `CustomPermissions` | `CustomPermissions` | no |
| `RequestId` | `string` | no |

## DescribeDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `VersionNumber` | `long` | no |
| `AliasName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Dashboard` | `Dashboard` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeDashboardDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `VersionNumber` | `long` | no |
| `AliasName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | no |
| `Errors` | `List<DashboardError>` | no |
| `Name` | `string` | no |
| `ResourceStatus` | `string` | no |
| `ThemeArn` | `string` | no |
| `Definition` | `DashboardVersionDefinition` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `DashboardPublishOptions` | `DashboardPublishOptions` | no |

## DescribeDashboardPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | no |
| `DashboardArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `LinkSharingConfiguration` | `LinkSharingConfiguration` | no |

## DescribeDashboardSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `SnapshotJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | no |
| `DashboardId` | `string` | no |
| `SnapshotJobId` | `string` | no |
| `UserConfiguration` | `SnapshotUserConfigurationRedacted` | no |
| `SnapshotConfiguration` | `SnapshotConfiguration` | no |
| `Arn` | `string` | no |
| `JobStatus` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDashboardSnapshotJobResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `SnapshotJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `JobStatus` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Result` | `SnapshotJobResult` | no |
| `ErrorInfo` | `SnapshotJobErrorInfo` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDashboardsQAConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardsQAStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSet` | `DataSet` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDataSetPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetArn` | `string` | no |
| `DataSetId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDataSetRefreshProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `DataSetRefreshProperties` | `DataSetRefreshProperties` | no |

## DescribeDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSource` | `DataSource` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDataSourcePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceArn` | `string` | no |
| `DataSourceId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeDefaultQBusinessApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `ApplicationId` | `string` | no |

## DescribeDlpSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DlpSettingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DlpSetting` | `DlpSettingDetails` | yes |
| `RequestId` | `string` | no |

## DescribeFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FlowId` | `string` | yes |
| `PublishState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Flow` | `FlowDetail` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Folder` | `Folder` | no |
| `RequestId` | `string` | no |

## DescribeFolderPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `Namespace` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `FolderId` | `string` | no |
| `Arn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `NextToken` | `string` | no |

## DescribeFolderResolvedPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `Namespace` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `FolderId` | `string` | no |
| `Arn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `NextToken` | `string` | no |

## DescribeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeGroupMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberName` | `string` | yes |
| `GroupName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupMember` | `GroupMember` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeIAMPolicyAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssignmentName` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IAMPolicyAssignment` | `IAMPolicyAssignment` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `IngestionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ingestion` | `Ingestion` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeIpRestriction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | no |
| `IpRestrictionRuleMap` | `Map<string>` | no |
| `VpcIdRestrictionRuleMap` | `Map<string>` | no |
| `VpcEndpointIdRestrictionRuleMap` | `Map<string>` | no |
| `Enabled` | `boolean` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeKeyRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DefaultKeyOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | no |
| `KeyRegistration` | `List<RegisteredCustomerManagedKey>` | no |
| `QDataKey` | `QDataKey` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBase` | `KnowledgeBase` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeKnowledgeBasePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseArn` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeLimitsProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `LimitsProfile` | yes |

## DescribeNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespace` | `NamespaceInfoV2` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeOAuthClientApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `OAuthClientApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OAuthClientApplication` | `OAuthClientApplication` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeQPersonalizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersonalizationMode` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeQuickSightQSearchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QSearchStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `ScheduleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshSchedule` | `RefreshSchedule` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `Arn` | `string` | no |

## DescribeRoleCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomPermissionsName` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeSelfUpgradeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelfUpgradeConfiguration` | `SelfUpgradeConfiguration` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |
| `MaxContributors` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `Space` | `SpaceDetails` | yes |
| `Contributors` | `List<SpaceContributor>` | no |
| `RequestId` | `string` | no |

## DescribeSpacePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |

## DescribeTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `VersionNumber` | `long` | no |
| `AliasName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Template` | `Template` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeTemplateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateAlias` | `TemplateAlias` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeTemplateDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `VersionNumber` | `long` | no |
| `AliasName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `TemplateId` | `string` | no |
| `Errors` | `List<TemplateError>` | no |
| `ResourceStatus` | `string` | no |
| `ThemeArn` | `string` | no |
| `Definition` | `TemplateVersionDefinition` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeTemplatePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateId` | `string` | no |
| `TemplateArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `VersionNumber` | `long` | no |
| `AliasName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Theme` | `Theme` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeThemeAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeAlias` | `ThemeAlias` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeThemePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeId` | `string` | no |
| `ThemeArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `Topic` | `TopicDetails` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `CustomInstructions` | `CustomInstructions` | no |

## DescribeTopicPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeTopicPermissionsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeTopicRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `RefreshId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshDetails` | `TopicRefreshDetails` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeTopicRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `DatasetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `DatasetArn` | `string` | no |
| `RefreshSchedule` | `TopicRefreshSchedule` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeTopicV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `Topic` | `TopicV2Details` | no |
| `CustomInstructions` | `CustomInstructions` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## DescribeVPCConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `VPCConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VPCConnection` | `VPCConnection` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## GenerateEmbedUrlForAnonymousUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SessionLifetimeInMinutes` | `long` | no |
| `Namespace` | `string` | yes |
| `SessionTags` | `List<SessionTag>` | no |
| `AuthorizedResourceArns` | `List<string>` | yes |
| `ExperienceConfiguration` | `AnonymousUserEmbeddingExperienceConfiguration` | yes |
| `AllowedDomains` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmbedUrl` | `string` | yes |
| `Status` | `integer` | yes |
| `RequestId` | `string` | yes |
| `AnonymousUserArn` | `string` | yes |

## GenerateEmbedUrlForRegisteredUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SessionLifetimeInMinutes` | `long` | no |
| `UserArn` | `string` | yes |
| `ExperienceConfiguration` | `RegisteredUserEmbeddingExperienceConfiguration` | yes |
| `AllowedDomains` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmbedUrl` | `string` | yes |
| `Status` | `integer` | yes |
| `RequestId` | `string` | yes |

## GenerateEmbedUrlForRegisteredUserWithIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SessionLifetimeInMinutes` | `long` | no |
| `ExperienceConfiguration` | `RegisteredUserEmbeddingExperienceConfiguration` | yes |
| `AllowedDomains` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmbedUrl` | `string` | yes |
| `Status` | `integer` | yes |
| `RequestId` | `string` | yes |

## GetDashboardEmbedUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `IdentityType` | `string` | yes |
| `SessionLifetimeInMinutes` | `long` | no |
| `UndoRedoDisabled` | `boolean` | no |
| `ResetDisabled` | `boolean` | no |
| `StatePersistenceEnabled` | `boolean` | no |
| `UserArn` | `string` | no |
| `Namespace` | `string` | no |
| `AdditionalDashboardIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmbedUrl` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## GetFlowMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `FlowId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `PublishState` | `string` | no |
| `UserCount` | `integer` | no |
| `RunCount` | `integer` | no |
| `CreatedTime` | `timestamp` | yes |
| `LastUpdatedTime` | `timestamp` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## GetFlowPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `FlowId` | `string` | yes |
| `Permissions` | `List<Permission>` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## GetIdentityContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `UserIdentifier` | `UserIdentifier` | yes |
| `Namespace` | `string` | no |
| `SessionExpiresAt` | `timestamp` | no |
| `ContextRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | yes |
| `RequestId` | `string` | yes |
| `Context` | `string` | no |

## GetSessionEmbedUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `EntryPoint` | `string` | no |
| `SessionLifetimeInMinutes` | `long` | no |
| `UserArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmbedUrl` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListActionConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionConnectorSummaries` | `List<ActionConnectorSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `AgentSummaries` | `List<AgentSummary>` | yes |
| `NextToken` | `string` | no |

## ListAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisSummaryList` | `List<AnalysisSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListApprovalPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `List<ApprovalPolicy>` | yes |
| `NextToken` | `string` | no |

## ListApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppSummaryList` | `List<AppSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListAssetBundleExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetBundleExportJobSummaryList` | `List<AssetBundleExportJobSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListAssetBundleImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetBundleImportJobSummaryList` | `List<AssetBundleImportJobSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListBrands

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Brands` | `List<BrandSummary>` | no |

## ListCustomPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `CustomPermissionsList` | `List<CustomPermissions>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListDashboardVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardVersionSummaryList` | `List<DashboardVersionSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardSummaryList` | `List<DashboardSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListDataSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetSummaries` | `List<DataSetSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSources` | `List<DataSource>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListDlpSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DlpSettingSummaries` | `List<DlpSettingSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowSummaryList` | `List<FlowSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListFolderMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `FolderMemberList` | `List<MemberIdArnPair>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListFolders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `FolderSummaryList` | `List<FolderSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListFoldersForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Folders` | `List<string>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListGroupMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupMemberList` | `List<GroupMember>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupList` | `List<Group>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListIAMPolicyAssignments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssignmentStatus` | `string` | no |
| `Namespace` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IAMPolicyAssignments` | `List<IAMPolicyAssignmentSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListIAMPolicyAssignmentsForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `UserName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActiveAssignments` | `List<ActiveIAMPolicyAssignment>` | no |
| `RequestId` | `string` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |

## ListIdentityPropagationConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Services` | `List<AuthorizedTargetsByService>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListIngestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `NextToken` | `string` | no |
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ingestions` | `List<Ingestion>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListKnowledgeBases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseSummaries` | `List<KnowledgeBaseSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListLimitsProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `resourceType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profiles` | `List<LimitsProfile>` | yes |
| `nextToken` | `string` | no |

## ListNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespaces` | `List<NamespaceInfoV2>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListOAuthClientApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OAuthClientApplications` | `List<OAuthClientApplicationSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListRefreshSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshSchedules` | `List<RefreshSchedule>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListRoleMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MembersList` | `List<string>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListSelfUpgrades

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelfUpgradeRequestDetails` | `List<SelfUpgradeRequestDetail>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListSpaceResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `SpaceResources` | `List<SpaceResourceSummary>` | yes |
| `RequestId` | `string` | no |

## ListSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `SpaceSummaries` | `List<SpaceSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListTemplateAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateAliasList` | `List<TemplateAlias>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `NextToken` | `string` | no |

## ListTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateVersionSummaryList` | `List<TemplateVersionSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateSummaryList` | `List<TemplateSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListThemeAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeAliasList` | `List<ThemeAlias>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `NextToken` | `string` | no |

## ListThemeVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeVersionSummaryList` | `List<ThemeVersionSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListThemes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeSummaryList` | `List<ThemeSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListTopicRefreshSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `RefreshSchedules` | `List<TopicRefreshScheduleSummary>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListTopicReviewedAnswers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `Answers` | `List<TopicReviewedAnswer>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## ListTopics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicsSummaries` | `List<TopicSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListTopicsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicSummaryList` | `List<TopicV2Summary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListUserGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupList` | `List<Group>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserList` | `List<User>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## ListUsersIndexCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `awsAccountId` | `string` | yes |
| `namespace` | `string` | no |
| `filters` | `List<UserIndexCapacityFilter>` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `users` | `List<UserIndexCapacity>` | no |
| `nextToken` | `string` | no |
| `requestId` | `string` | no |

## ListVPCConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VPCConnectionSummaries` | `List<VPCConnectionSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## PredictQAResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `QueryText` | `string` | yes |
| `IncludeQuickSightQIndex` | `string` | no |
| `IncludeGeneratedAnswer` | `string` | no |
| `MaxTopicsToConsider` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrimaryResult` | `QAResult` | no |
| `AdditionalResults` | `List<QAResult>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## PutDataSetRefreshProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `DataSetRefreshProperties` | `DataSetRefreshProperties` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## RegisterUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityType` | `string` | yes |
| `Email` | `string` | yes |
| `UserRole` | `string` | yes |
| `IamArn` | `string` | no |
| `SessionName` | `string` | no |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `UserName` | `string` | no |
| `CustomPermissionsName` | `string` | no |
| `ExternalLoginFederationProviderType` | `string` | no |
| `CustomFederationProviderUrl` | `string` | no |
| `ExternalLoginId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |
| `UserInvitationUrl` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## RestoreAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |
| `RestoreToFolders` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `AnalysisId` | `string` | no |
| `RequestId` | `string` | no |
| `RestorationFailedFolderArns` | `List<string>` | no |

## SearchActionConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<ActionConnectorSearchFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `ActionConnectorSummaries` | `List<ActionConnectorSummary>` | no |

## SearchAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<AgentSearchFilter>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentSummaries` | `List<AgentSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## SearchAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<AnalysisSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisSummaryList` | `List<AnalysisSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## SearchApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<SearchAppsFilter>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppSummaryList` | `List<AppSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## SearchDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<DashboardSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardSummaryList` | `List<DashboardSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## SearchDataSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<DataSetSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetSummaries` | `List<DataSetSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## SearchDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<DataSourceSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceSummaries` | `List<DataSourceSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## SearchFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<SearchFlowsFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowSummaryList` | `List<FlowSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## SearchFolders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<FolderSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `FolderSummaryList` | `List<FolderSummary>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## SearchGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Namespace` | `string` | yes |
| `Filters` | `List<GroupSearchFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupList` | `List<Group>` | no |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## SearchKnowledgeBases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<KnowledgeBaseSearchFilter>` | no |
| `SortBy` | `KnowledgeBaseSortBy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseSummaries` | `List<KnowledgeBaseSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## SearchSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<SpaceQuicksightSearchFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `SpaceSummaries` | `List<SpaceSummary>` | yes |
| `NextToken` | `string` | no |
| `RequestId` | `string` | no |

## SearchTopics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<TopicSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicSummaryList` | `List<TopicSummary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## SearchTopicsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Filters` | `List<TopicSearchFilter>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicSummaryList` | `List<TopicV2Summary>` | no |
| `NextToken` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## StartAssetBundleExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssetBundleExportJobId` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |
| `IncludeAllDependencies` | `boolean` | no |
| `ExportFormat` | `string` | yes |
| `CloudFormationOverridePropertyConfiguration` | `AssetBundleCloudFormationOverridePropertyConfiguration` | no |
| `IncludePermissions` | `boolean` | no |
| `IncludeTags` | `boolean` | no |
| `ValidationStrategy` | `AssetBundleExportJobValidationStrategy` | no |
| `IncludeFolderMemberships` | `boolean` | no |
| `IncludeFolderMembers` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetBundleExportJobId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## StartAssetBundleImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssetBundleImportJobId` | `string` | yes |
| `AssetBundleImportSource` | `AssetBundleImportSource` | yes |
| `OverrideParameters` | `AssetBundleImportJobOverrideParameters` | no |
| `FailureAction` | `string` | no |
| `OverridePermissions` | `AssetBundleImportJobOverridePermissions` | no |
| `OverrideTags` | `AssetBundleImportJobOverrideTags` | no |
| `OverrideValidationStrategy` | `AssetBundleImportJobOverrideValidationStrategy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetBundleImportJobId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## StartAutomationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AutomationGroupId` | `string` | yes |
| `AutomationId` | `string` | yes |
| `InputPayload` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `JobId` | `string` | yes |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## StartDashboardSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `SnapshotJobId` | `string` | yes |
| `UserConfiguration` | `SnapshotUserConfiguration` | no |
| `SnapshotConfiguration` | `SnapshotConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `SnapshotJobId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## StartDashboardSnapshotJobSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `ScheduleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateAccountCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomPermissionsName` | `string` | yes |
| `AwsAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateAccountCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |
| `AccountCustomization` | `AccountCustomization` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AwsAccountId` | `string` | no |
| `Namespace` | `string` | no |
| `AccountCustomization` | `AccountCustomization` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DefaultNamespace` | `string` | yes |
| `NotificationEmail` | `string` | no |
| `TerminationProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateActionConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ActionConnectorId` | `string` | yes |
| `Name` | `string` | yes |
| `AuthenticationConfig` | `AuthConfig` | yes |
| `Description` | `string` | no |
| `VpcConnectionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ActionConnectorId` | `string` | no |
| `RequestId` | `string` | no |
| `UpdateStatus` | `string` | no |
| `Status` | `integer` | no |

## UpdateActionConnectorPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ActionConnectorId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ActionConnectorId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `Permissions` | `List<ResourcePermission>` | no |

## UpdateAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `IconId` | `string` | no |
| `StarterPrompts` | `List<string>` | no |
| `WelcomeMessage` | `string` | no |
| `CustomPromptInput` | `CustomPromptInput` | no |
| `SpacesToAdd` | `List<string>` | no |
| `SpacesToRemove` | `List<string>` | no |
| `ActionConnectorsToAdd` | `List<string>` | no |
| `ActionConnectorsToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `AgentId` | `string` | yes |
| `AgentStatus` | `string` | yes |
| `FailedToAddSpaces` | `List<FailedToUpdateAssociation>` | no |
| `FailedToRemoveSpaces` | `List<FailedToUpdateAssociation>` | no |
| `FailedToAddActionConnectors` | `List<FailedToUpdateAssociation>` | no |
| `FailedToRemoveActionConnectors` | `List<FailedToUpdateAssociation>` | no |
| `RequestId` | `string` | no |

## UpdateAgentPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `AgentId` | `string` | yes |
| `RequestId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |

## UpdateAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |
| `Name` | `string` | yes |
| `Parameters` | `Parameters` | no |
| `SourceEntity` | `AnalysisSourceEntity` | no |
| `ThemeArn` | `string` | no |
| `Definition` | `AnalysisDefinition` | no |
| `ValidationStrategy` | `ValidationStrategy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AnalysisId` | `string` | no |
| `UpdateStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateAnalysisPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AnalysisId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisArn` | `string` | no |
| `AnalysisId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateAppPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AppId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |
| `Visibility` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AppId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Visibility` | `string` | no |
| `RequestId` | `string` | no |

## UpdateApplicationWithTokenExchangeGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateApprovalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Actions` | `List<string>` | no |
| `AssetTypes` | `List<string>` | no |
| `ApplicableTo` | `ApplicableTo` | no |
| `ApprovalGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `ApprovalPolicy` | yes |

## UpdateBrand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandId` | `string` | yes |
| `BrandDefinition` | `BrandDefinition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `BrandDetail` | `BrandDetail` | no |
| `BrandDefinition` | `BrandDefinition` | no |

## UpdateBrandAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `BrandArn` | `string` | no |

## UpdateBrandPublishedVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `BrandId` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `VersionId` | `string` | no |

## UpdateCustomPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `CustomPermissionsName` | `string` | yes |
| `Capabilities` | `Capabilities` | no |
| `Governance` | `Governance` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `RequestId` | `string` | no |

## UpdateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `Name` | `string` | yes |
| `SourceEntity` | `DashboardSourceEntity` | no |
| `Parameters` | `Parameters` | no |
| `VersionDescription` | `string` | no |
| `DashboardPublishOptions` | `DashboardPublishOptions` | no |
| `ThemeArn` | `string` | no |
| `Definition` | `DashboardVersionDefinition` | no |
| `ValidationStrategy` | `ValidationStrategy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VersionArn` | `string` | no |
| `DashboardId` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateDashboardLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `LinkEntities` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `DashboardArn` | `string` | no |
| `LinkEntities` | `List<string>` | no |

## UpdateDashboardPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |
| `GrantLinkPermissions` | `List<ResourcePermission>` | no |
| `RevokeLinkPermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardArn` | `string` | no |
| `DashboardId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |
| `LinkSharingConfiguration` | `LinkSharingConfiguration` | no |

## UpdateDashboardPublishedVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardId` | `string` | yes |
| `VersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | no |
| `DashboardArn` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateDashboardsQAConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DashboardsQAStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardsQAStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `Name` | `string` | yes |
| `PhysicalTableMap` | `Map<PhysicalTable>` | yes |
| `LogicalTableMap` | `Map<LogicalTable>` | no |
| `ImportMode` | `string` | yes |
| `ColumnGroups` | `List<ColumnGroup>` | no |
| `FieldFolders` | `Map<FieldFolder>` | no |
| `RowLevelPermissionDataSet` | `RowLevelPermissionDataSet` | no |
| `RowLevelPermissionTagConfiguration` | `RowLevelPermissionTagConfiguration` | no |
| `ColumnLevelPermissionRules` | `List<ColumnLevelPermissionRule>` | no |
| `DataSetUsageConfiguration` | `DataSetUsageConfiguration` | no |
| `DatasetParameters` | `List<DatasetParameter>` | no |
| `PerformanceConfiguration` | `PerformanceConfiguration` | no |
| `DataPrepConfiguration` | `DataPrepConfiguration` | no |
| `SemanticModelConfiguration` | `SemanticModelConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `DataSetId` | `string` | no |
| `IngestionArn` | `string` | no |
| `IngestionId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateDataSetPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetArn` | `string` | no |
| `DataSetId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSourceId` | `string` | yes |
| `Name` | `string` | yes |
| `DataSourceParameters` | `DataSourceParameters` | no |
| `Credentials` | `DataSourceCredentials` | no |
| `VpcConnectionProperties` | `VpcConnectionProperties` | no |
| `SslProperties` | `SslProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `DataSourceId` | `string` | no |
| `UpdateStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateDataSourcePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DataSourceId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceArn` | `string` | no |
| `DataSourceId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateDefaultQBusinessApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | no |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateDlpSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `DlpSettingId` | `string` | yes |
| `Name` | `string` | no |
| `ProviderType` | `string` | no |
| `ProviderConfig` | `ProviderConfig` | no |
| `ProviderOutageAction` | `string` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `DlpSettingId` | `string` | yes |
| `RequestId` | `string` | no |

## UpdateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FlowId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `FlowDefinition` | `SensitiveDocument` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `FlowId` | `string` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateFlowPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FlowId` | `string` | yes |
| `GrantPermissions` | `List<Permission>` | no |
| `RevokePermissions` | `List<Permission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | yes |
| `Permissions` | `List<Permission>` | yes |
| `RequestId` | `string` | yes |
| `FlowId` | `string` | yes |

## UpdateFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `FolderId` | `string` | no |
| `RequestId` | `string` | no |

## UpdateFolderPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `FolderId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `Arn` | `string` | no |
| `FolderId` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |

## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `Description` | `string` | no |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateIAMPolicyAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `AssignmentName` | `string` | yes |
| `Namespace` | `string` | yes |
| `AssignmentStatus` | `string` | no |
| `PolicyArn` | `string` | no |
| `Identities` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentName` | `string` | no |
| `AssignmentId` | `string` | no |
| `PolicyArn` | `string` | no |
| `Identities` | `Map<List<string>>` | no |
| `AssignmentStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateIdentityPropagationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Service` | `string` | yes |
| `AuthorizedTargets` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateIpRestriction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `IpRestrictionRuleMap` | `Map<string>` | no |
| `VpcIdRestrictionRuleMap` | `Map<string>` | no |
| `VpcEndpointIdRestrictionRuleMap` | `Map<string>` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateKeyRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KeyRegistration` | `List<RegisteredCustomerManagedKey>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedKeyRegistration` | `List<FailedKeyRegistrationEntry>` | no |
| `SuccessfulKeyRegistration` | `List<SuccessfulKeyRegistrationEntry>` | no |
| `RequestId` | `string` | no |

## UpdateKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `KnowledgeBaseConfiguration` | `KnowledgeBaseConfiguration` | no |
| `MediaExtractionConfiguration` | `MediaExtractionConfiguration` | no |
| `IsEmailNotificationOptedForIngestionFailures` | `boolean` | no |
| `AccessControlConfiguration` | `AccessControlConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseArn` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateKnowledgeBasePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KnowledgeBaseArn` | `string` | yes |
| `KnowledgeBaseId` | `string` | yes |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateLimitsProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `accountId` | `string` | yes |
| `profileName` | `string` | no |
| `description` | `string` | no |
| `resourceLimits` | `Map<ProfileLimitValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## UpdateOAuthClientApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `OAuthClientApplicationId` | `string` | yes |
| `Name` | `string` | yes |
| `ClientId` | `string` | no |
| `ClientSecret` | `string` | no |
| `OAuthTokenEndpointUrl` | `string` | no |
| `OAuthAuthorizationEndpointUrl` | `string` | no |
| `OAuthScopes` | `string` | no |
| `DataSourceType` | `string` | no |
| `IdentityProviderVpcConnectionProperties` | `VpcConnectionProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `OAuthClientApplicationId` | `string` | no |
| `UpdateStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdatePublicSharingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `PublicSharingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateQPersonalizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `PersonalizationMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersonalizationMode` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateQuickSightQSearchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `QSearchStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QSearchStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Schedule` | `RefreshSchedule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |
| `RequestId` | `string` | no |
| `ScheduleId` | `string` | no |
| `Arn` | `string` | no |

## UpdateRoleCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomPermissionsName` | `string` | yes |
| `Role` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateSPICECapacityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `PurchaseMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateSelfUpgrade

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `UpgradeRequestId` | `string` | yes |
| `Action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelfUpgradeRequestDetail` | `SelfUpgradeRequestDetail` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateSelfUpgradeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `SelfUpgradeStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `RequestId` | `string` | no |

## UpdateSpacePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `permissions` | `List<ResourcePermission>` | no |
| `requestId` | `string` | no |

## UpdateSpaceResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `SpaceId` | `string` | yes |
| `AddResources` | `List<SpaceResourceOperation>` | no |
| `RemoveResources` | `List<SpaceResourceOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `spaceArn` | `string` | no |
| `FailedResourceOperations` | `List<FailedSpaceResourceOperation>` | no |
| `RequestId` | `string` | no |

## UpdateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `SourceEntity` | `TemplateSourceEntity` | no |
| `VersionDescription` | `string` | no |
| `Name` | `string` | no |
| `Definition` | `TemplateVersionDefinition` | no |
| `ValidationStrategy` | `ValidationStrategy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateId` | `string` | no |
| `Arn` | `string` | no |
| `VersionArn` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateTemplateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `AliasName` | `string` | yes |
| `TemplateVersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateAlias` | `TemplateAlias` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateTemplatePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TemplateId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateId` | `string` | no |
| `TemplateArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `Name` | `string` | no |
| `BaseThemeId` | `string` | yes |
| `VersionDescription` | `string` | no |
| `Configuration` | `ThemeConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeId` | `string` | no |
| `Arn` | `string` | no |
| `VersionArn` | `string` | no |
| `CreationStatus` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateThemeAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `AliasName` | `string` | yes |
| `ThemeVersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeAlias` | `ThemeAlias` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateThemePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `ThemeId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThemeId` | `string` | no |
| `ThemeArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `Topic` | `TopicDetails` | yes |
| `CustomInstructions` | `CustomInstructions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `Arn` | `string` | no |
| `RefreshArn` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateTopicPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateTopicPermissionsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `GrantPermissions` | `List<ResourcePermission>` | no |
| `RevokePermissions` | `List<ResourcePermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `Permissions` | `List<ResourcePermission>` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateTopicRefreshSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `DatasetId` | `string` | yes |
| `RefreshSchedule` | `TopicRefreshSchedule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicId` | `string` | no |
| `TopicArn` | `string` | no |
| `DatasetArn` | `string` | no |
| `Status` | `integer` | no |
| `RequestId` | `string` | no |

## UpdateTopicV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `TopicId` | `string` | yes |
| `Topic` | `TopicV2Details` | yes |
| `CustomInstructions` | `CustomInstructions` | no |
| `PublishOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `TopicId` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `Email` | `string` | yes |
| `Role` | `string` | yes |
| `CustomPermissionsName` | `string` | no |
| `UnapplyCustomPermissions` | `boolean` | no |
| `ExternalLoginFederationProviderType` | `string` | no |
| `CustomFederationProviderUrl` | `string` | no |
| `ExternalLoginId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateUserCustomPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AwsAccountId` | `string` | yes |
| `Namespace` | `string` | yes |
| `CustomPermissionsName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

## UpdateVPCConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `VPCConnectionId` | `string` | yes |
| `Name` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `SecurityGroupIds` | `List<string>` | yes |
| `DnsResolvers` | `List<string>` | no |
| `RoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `VPCConnectionId` | `string` | no |
| `UpdateStatus` | `string` | no |
| `AvailabilityStatus` | `string` | no |
| `RequestId` | `string` | no |
| `Status` | `integer` | no |

