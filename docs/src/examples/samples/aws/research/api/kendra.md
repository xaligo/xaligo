# AWSKendraFrontendService

API version: 2019-02-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kendra/2019-02-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateEntitiesToExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `EntityList` | `List<EntityConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntityList` | `List<FailedEntity>` | no |

## AssociatePersonasToEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `Personas` | `List<EntityPersonaConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntityList` | `List<FailedEntity>` | no |

## BatchDeleteDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `DocumentIdList` | `List<string>` | yes |
| `DataSourceSyncJobMetricTarget` | `DataSourceSyncJobMetricTarget` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedDocuments` | `List<BatchDeleteDocumentResponseFailedDocument>` | no |

## BatchDeleteFeaturedResultsSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `FeaturedResultsSetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchDeleteFeaturedResultsSetError>` | yes |

## BatchGetDocumentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `DocumentInfoList` | `List<DocumentInfo>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchGetDocumentStatusResponseError>` | no |
| `DocumentStatusList` | `List<Status>` | no |

## BatchPutDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `RoleArn` | `string` | no |
| `Documents` | `List<Document>` | yes |
| `CustomDocumentEnrichmentConfiguration` | `CustomDocumentEnrichmentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedDocuments` | `List<BatchPutDocumentResponseFailedDocument>` | no |

## ClearQuerySuggestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAccessControlConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `AccessControlList` | `List<Principal>` | no |
| `HierarchicalAccessControlList` | `List<HierarchicalPrincipal>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

## CreateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IndexId` | `string` | yes |
| `Type` | `string` | yes |
| `Configuration` | `DataSourceConfiguration` | no |
| `VpcConfiguration` | `DataSourceVpcConfiguration` | no |
| `Description` | `string` | no |
| `Schedule` | `string` | no |
| `RoleArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |
| `LanguageCode` | `string` | no |
| `CustomDocumentEnrichmentConfiguration` | `CustomDocumentEnrichmentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

## CreateExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IndexId` | `string` | yes |
| `RoleArn` | `string` | no |
| `Configuration` | `ExperienceConfiguration` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

## CreateFaq

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `S3Path` | `S3Path` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `FileFormat` | `string` | no |
| `ClientToken` | `string` | no |
| `LanguageCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## CreateFeaturedResultsSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `FeaturedResultsSetName` | `string` | yes |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `Status` | `string` | no |
| `QueryTexts` | `List<string>` | no |
| `FeaturedDocuments` | `List<FeaturedDocument>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeaturedResultsSet` | `FeaturedResultsSet` | no |

## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Edition` | `string` | no |
| `RoleArn` | `string` | yes |
| `ServerSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `UserTokenConfigurations` | `List<UserTokenConfiguration>` | no |
| `UserContextPolicy` | `string` | no |
| `UserGroupResolutionConfiguration` | `UserGroupResolutionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## CreateQuerySuggestionsBlockList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `SourceS3Path` | `S3Path` | yes |
| `ClientToken` | `string` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## CreateThesaurus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `SourceS3Path` | `S3Path` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## DeleteAccessControlConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFaq

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePrincipalMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `DataSourceId` | `string` | no |
| `GroupId` | `string` | yes |
| `OrderingId` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQuerySuggestionsBlockList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThesaurus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccessControlConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ErrorMessage` | `string` | no |
| `AccessControlList` | `List<Principal>` | no |
| `HierarchicalAccessControlList` | `List<HierarchicalPrincipal>` | no |

## DescribeDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `IndexId` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `Configuration` | `DataSourceConfiguration` | no |
| `VpcConfiguration` | `DataSourceVpcConfiguration` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `Schedule` | `string` | no |
| `RoleArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `LanguageCode` | `string` | no |
| `CustomDocumentEnrichmentConfiguration` | `CustomDocumentEnrichmentConfiguration` | no |

## DescribeExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `IndexId` | `string` | no |
| `Name` | `string` | no |
| `Endpoints` | `List<ExperienceEndpoint>` | no |
| `Configuration` | `ExperienceConfiguration` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `RoleArn` | `string` | no |
| `ErrorMessage` | `string` | no |

## DescribeFaq

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `IndexId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `S3Path` | `S3Path` | no |
| `Status` | `string` | no |
| `RoleArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `FileFormat` | `string` | no |
| `LanguageCode` | `string` | no |

## DescribeFeaturedResultsSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `FeaturedResultsSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeaturedResultsSetId` | `string` | no |
| `FeaturedResultsSetName` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `QueryTexts` | `List<string>` | no |
| `FeaturedDocumentsWithMetadata` | `List<FeaturedDocumentWithMetadata>` | no |
| `FeaturedDocumentsMissing` | `List<FeaturedDocumentMissing>` | no |
| `LastUpdatedTimestamp` | `long` | no |
| `CreationTimestamp` | `long` | no |

## DescribeIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Id` | `string` | no |
| `Edition` | `string` | no |
| `RoleArn` | `string` | no |
| `ServerSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `Status` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `DocumentMetadataConfigurations` | `List<DocumentMetadataConfiguration>` | no |
| `IndexStatistics` | `IndexStatistics` | no |
| `ErrorMessage` | `string` | no |
| `CapacityUnits` | `CapacityUnitsConfiguration` | no |
| `UserTokenConfigurations` | `List<UserTokenConfiguration>` | no |
| `UserContextPolicy` | `string` | no |
| `UserGroupResolutionConfiguration` | `UserGroupResolutionConfiguration` | no |

## DescribePrincipalMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `DataSourceId` | `string` | no |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | no |
| `DataSourceId` | `string` | no |
| `GroupId` | `string` | no |
| `GroupOrderingIdSummaries` | `List<GroupOrderingIdSummary>` | no |

## DescribeQuerySuggestionsBlockList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `ErrorMessage` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `SourceS3Path` | `S3Path` | no |
| `ItemCount` | `integer` | no |
| `FileSizeBytes` | `long` | no |
| `RoleArn` | `string` | no |

## DescribeQuerySuggestionsConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mode` | `string` | no |
| `Status` | `string` | no |
| `QueryLogLookBackWindowInDays` | `integer` | no |
| `IncludeQueriesWithoutUserInformation` | `boolean` | no |
| `MinimumNumberOfQueryingUsers` | `integer` | no |
| `MinimumQueryCount` | `integer` | no |
| `LastSuggestionsBuildTime` | `timestamp` | no |
| `LastClearTime` | `timestamp` | no |
| `TotalSuggestionsCount` | `integer` | no |
| `AttributeSuggestionsConfig` | `AttributeSuggestionsDescribeConfig` | no |

## DescribeThesaurus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `IndexId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `ErrorMessage` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `RoleArn` | `string` | no |
| `SourceS3Path` | `S3Path` | no |
| `FileSizeBytes` | `long` | no |
| `TermCount` | `long` | no |
| `SynonymRuleCount` | `long` | no |

## DisassociateEntitiesFromExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `EntityList` | `List<EntityConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntityList` | `List<FailedEntity>` | no |

## DisassociatePersonasFromEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `EntityIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntityList` | `List<FailedEntity>` | no |

## GetQuerySuggestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `QueryText` | `string` | yes |
| `MaxSuggestionsCount` | `integer` | no |
| `SuggestionTypes` | `List<string>` | no |
| `AttributeSuggestionsConfig` | `AttributeSuggestionsGetConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuerySuggestionsId` | `string` | no |
| `Suggestions` | `List<Suggestion>` | no |

## GetSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Interval` | `string` | yes |
| `MetricType` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapShotTimeFilter` | `TimeRange` | no |
| `SnapshotsDataHeader` | `List<string>` | no |
| `SnapshotsData` | `List<List<string>>` | no |
| `NextToken` | `string` | no |

## ListAccessControlConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AccessControlConfigurations` | `List<AccessControlConfigurationSummary>` | yes |

## ListDataSourceSyncJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StartTimeFilter` | `TimeRange` | no |
| `StatusFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `History` | `List<DataSourceSyncJob>` | no |
| `NextToken` | `string` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryItems` | `List<DataSourceSummary>` | no |
| `NextToken` | `string` | no |

## ListEntityPersonas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryItems` | `List<PersonasSummary>` | no |
| `NextToken` | `string` | no |

## ListExperienceEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryItems` | `List<ExperienceEntitiesSummary>` | no |
| `NextToken` | `string` | no |

## ListExperiences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryItems` | `List<ExperiencesSummary>` | no |
| `NextToken` | `string` | no |

## ListFaqs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FaqSummaryItems` | `List<FaqSummary>` | no |

## ListFeaturedResultsSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeaturedResultsSetSummaryItems` | `List<FeaturedResultsSetSummary>` | no |
| `NextToken` | `string` | no |

## ListGroupsOlderThanOrderingId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `DataSourceId` | `string` | no |
| `OrderingId` | `long` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupsSummaries` | `List<GroupSummary>` | no |
| `NextToken` | `string` | no |

## ListIndices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexConfigurationSummaryItems` | `List<IndexConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListQuerySuggestionsBlockLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlockListSummaryItems` | `List<QuerySuggestionsBlockListSummary>` | no |
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

## ListThesauri

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ThesaurusSummaryItems` | `List<ThesaurusSummary>` | no |

## PutPrincipalMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `DataSourceId` | `string` | no |
| `GroupId` | `string` | yes |
| `GroupMembers` | `GroupMembers` | yes |
| `OrderingId` | `long` | no |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Query

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `QueryText` | `string` | no |
| `AttributeFilter` | `AttributeFilter` | no |
| `Facets` | `List<Facet>` | no |
| `RequestedDocumentAttributes` | `List<string>` | no |
| `QueryResultTypeFilter` | `string` | no |
| `DocumentRelevanceOverrideConfigurations` | `List<DocumentRelevanceConfiguration>` | no |
| `PageNumber` | `integer` | no |
| `PageSize` | `integer` | no |
| `SortingConfiguration` | `SortingConfiguration` | no |
| `SortingConfigurations` | `List<SortingConfiguration>` | no |
| `UserContext` | `UserContext` | no |
| `VisitorId` | `string` | no |
| `SpellCorrectionConfiguration` | `SpellCorrectionConfiguration` | no |
| `CollapseConfiguration` | `CollapseConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | no |
| `ResultItems` | `List<QueryResultItem>` | no |
| `FacetResults` | `List<FacetResult>` | no |
| `TotalNumberOfResults` | `integer` | no |
| `Warnings` | `List<Warning>` | no |
| `SpellCorrectedQueries` | `List<SpellCorrectedQuery>` | no |
| `FeaturedResultsItems` | `List<FeaturedResultsItem>` | no |

## Retrieve

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `QueryText` | `string` | yes |
| `AttributeFilter` | `AttributeFilter` | no |
| `RequestedDocumentAttributes` | `List<string>` | no |
| `DocumentRelevanceOverrideConfigurations` | `List<DocumentRelevanceConfiguration>` | no |
| `PageNumber` | `integer` | no |
| `PageSize` | `integer` | no |
| `UserContext` | `UserContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | no |
| `ResultItems` | `List<RetrieveResultItem>` | no |

## StartDataSourceSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionId` | `string` | no |

## StopDataSourceSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IndexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubmitFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `QueryId` | `string` | yes |
| `ClickFeedbackItems` | `List<ClickFeedback>` | no |
| `RelevanceFeedbackItems` | `List<RelevanceFeedback>` | no |

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


## UpdateAccessControlConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `AccessControlList` | `List<Principal>` | no |
| `HierarchicalAccessControlList` | `List<HierarchicalPrincipal>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `IndexId` | `string` | yes |
| `Configuration` | `DataSourceConfiguration` | no |
| `VpcConfiguration` | `DataSourceVpcConfiguration` | no |
| `Description` | `string` | no |
| `Schedule` | `string` | no |
| `RoleArn` | `string` | no |
| `LanguageCode` | `string` | no |
| `CustomDocumentEnrichmentConfiguration` | `CustomDocumentEnrichmentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `IndexId` | `string` | yes |
| `RoleArn` | `string` | no |
| `Configuration` | `ExperienceConfiguration` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFeaturedResultsSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `FeaturedResultsSetId` | `string` | yes |
| `FeaturedResultsSetName` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `QueryTexts` | `List<string>` | no |
| `FeaturedDocuments` | `List<FeaturedDocument>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeaturedResultsSet` | `FeaturedResultsSet` | no |

## UpdateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `RoleArn` | `string` | no |
| `Description` | `string` | no |
| `DocumentMetadataConfigurationUpdates` | `List<DocumentMetadataConfiguration>` | no |
| `CapacityUnits` | `CapacityUnitsConfiguration` | no |
| `UserTokenConfigurations` | `List<UserTokenConfiguration>` | no |
| `UserContextPolicy` | `string` | no |
| `UserGroupResolutionConfiguration` | `UserGroupResolutionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQuerySuggestionsBlockList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `SourceS3Path` | `S3Path` | no |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQuerySuggestionsConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexId` | `string` | yes |
| `Mode` | `string` | no |
| `QueryLogLookBackWindowInDays` | `integer` | no |
| `IncludeQueriesWithoutUserInformation` | `boolean` | no |
| `MinimumNumberOfQueryingUsers` | `integer` | no |
| `MinimumQueryCount` | `integer` | no |
| `AttributeSuggestionsConfig` | `AttributeSuggestionsUpdateConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateThesaurus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `IndexId` | `string` | yes |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `SourceS3Path` | `S3Path` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


