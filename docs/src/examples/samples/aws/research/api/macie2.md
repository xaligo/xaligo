# Amazon Macie 2

API version: 2020-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/macie2/2020-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `administratorAccountId` | `string` | no |
| `invitationId` | `string` | yes |
| `masterAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchGetCustomDataIdentifiers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDataIdentifiers` | `List<BatchGetCustomDataIdentifierSummary>` | no |
| `notFoundIdentifierIds` | `List<string>` | no |

## BatchUpdateAutomatedDiscoveryAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<AutomatedDiscoveryAccountUpdate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<AutomatedDiscoveryAccountUpdateError>` | no |

## CreateAllowList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `criteria` | `AllowListCriteria` | yes |
| `description` | `string` | no |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |

## CreateClassificationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `allowListIds` | `List<string>` | no |
| `clientToken` | `string` | yes |
| `customDataIdentifierIds` | `List<string>` | no |
| `description` | `string` | no |
| `initialRun` | `boolean` | no |
| `jobType` | `string` | yes |
| `managedDataIdentifierIds` | `List<string>` | no |
| `managedDataIdentifierSelector` | `string` | no |
| `name` | `string` | yes |
| `s3JobDefinition` | `S3JobDefinition` | yes |
| `samplingPercentage` | `integer` | no |
| `scheduleFrequency` | `JobScheduleFrequency` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobId` | `string` | no |

## CreateCustomDataIdentifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `ignoreWords` | `List<string>` | no |
| `keywords` | `List<string>` | no |
| `maximumMatchDistance` | `integer` | no |
| `name` | `string` | yes |
| `regex` | `string` | yes |
| `severityLevels` | `List<SeverityLevel>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDataIdentifierId` | `string` | no |

## CreateFindingsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `findingCriteria` | `FindingCriteria` | yes |
| `name` | `string` | yes |
| `position` | `integer` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |

## CreateInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |
| `disableEmailNotification` | `boolean` | no |
| `message` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `unprocessedAccounts` | `List<UnprocessedAccount>` | no |

## CreateMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `account` | `AccountDetail` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |

## CreateSampleFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeclineInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `unprocessedAccounts` | `List<UnprocessedAccount>` | no |

## DeleteAllowList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `ignoreJobChecks` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomDataIdentifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFindingsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `unprocessedAccounts` | `List<UnprocessedAccount>` | no |

## DeleteMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeBuckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `criteria` | `Map<BucketCriteriaAdditionalProperties>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortCriteria` | `BucketSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buckets` | `List<BucketMetadata>` | no |
| `nextToken` | `string` | no |

## DescribeClassificationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `allowListIds` | `List<string>` | no |
| `clientToken` | `string` | no |
| `createdAt` | `timestamp` | no |
| `customDataIdentifierIds` | `List<string>` | no |
| `description` | `string` | no |
| `initialRun` | `boolean` | no |
| `jobArn` | `string` | no |
| `jobId` | `string` | no |
| `jobStatus` | `string` | no |
| `jobType` | `string` | no |
| `lastRunErrorStatus` | `LastRunErrorStatus` | no |
| `lastRunTime` | `timestamp` | no |
| `managedDataIdentifierIds` | `List<string>` | no |
| `managedDataIdentifierSelector` | `string` | no |
| `name` | `string` | no |
| `s3JobDefinition` | `S3JobDefinition` | no |
| `samplingPercentage` | `integer` | no |
| `scheduleFrequency` | `JobScheduleFrequency` | no |
| `statistics` | `Statistics` | no |
| `tags` | `Map<string>` | no |
| `userPausedDetails` | `UserPausedDetails` | no |

## DescribeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnable` | `boolean` | no |
| `maxAccountLimitReached` | `boolean` | no |

## DisableMacie

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFromAdministratorAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFromMasterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableMacie

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `findingPublishingFrequency` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccountId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAdministratorAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `administrator` | `Invitation` | no |

## GetAllowList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `criteria` | `AllowListCriteria` | no |
| `description` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `status` | `AllowListStatus` | no |
| `tags` | `Map<string>` | no |
| `updatedAt` | `timestamp` | no |

## GetAutomatedDiscoveryConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnableOrganizationMembers` | `string` | no |
| `classificationScopeId` | `string` | no |
| `disabledAt` | `timestamp` | no |
| `firstEnabledAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `sensitivityInspectionTemplateId` | `string` | no |
| `status` | `string` | no |

## GetBucketStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketCount` | `long` | no |
| `bucketCountByEffectivePermission` | `BucketCountByEffectivePermission` | no |
| `bucketCountByEncryptionType` | `BucketCountByEncryptionType` | no |
| `bucketCountByObjectEncryptionRequirement` | `BucketCountPolicyAllowsUnencryptedObjectUploads` | no |
| `bucketCountBySharedAccessType` | `BucketCountBySharedAccessType` | no |
| `bucketStatisticsBySensitivity` | `BucketStatisticsBySensitivity` | no |
| `classifiableObjectCount` | `long` | no |
| `classifiableSizeInBytes` | `long` | no |
| `lastUpdated` | `timestamp` | no |
| `objectCount` | `long` | no |
| `sizeInBytes` | `long` | no |
| `sizeInBytesCompressed` | `long` | no |
| `unclassifiableObjectCount` | `ObjectLevelStatistics` | no |
| `unclassifiableObjectSizeInBytes` | `ObjectLevelStatistics` | no |

## GetClassificationExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `ClassificationExportConfiguration` | no |

## GetClassificationScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `s3` | `S3ClassificationScope` | no |

## GetCustomDataIdentifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `deleted` | `boolean` | no |
| `description` | `string` | no |
| `id` | `string` | no |
| `ignoreWords` | `List<string>` | no |
| `keywords` | `List<string>` | no |
| `maximumMatchDistance` | `integer` | no |
| `name` | `string` | no |
| `regex` | `string` | no |
| `severityLevels` | `List<SeverityLevel>` | no |
| `tags` | `Map<string>` | no |

## GetFindingStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingCriteria` | `FindingCriteria` | no |
| `groupBy` | `string` | yes |
| `size` | `integer` | no |
| `sortCriteria` | `FindingStatisticsSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `countsByGroup` | `List<GroupCount>` | no |

## GetFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingIds` | `List<string>` | yes |
| `sortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<Finding>` | no |

## GetFindingsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | no |
| `arn` | `string` | no |
| `description` | `string` | no |
| `findingCriteria` | `FindingCriteria` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `position` | `integer` | no |
| `tags` | `Map<string>` | no |

## GetFindingsPublicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityHubConfiguration` | `SecurityHubConfiguration` | no |

## GetInvitationsCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invitationsCount` | `long` | no |

## GetMacieSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAt` | `timestamp` | no |
| `findingPublishingFrequency` | `string` | no |
| `serviceRole` | `string` | no |
| `status` | `string` | no |
| `updatedAt` | `timestamp` | no |

## GetMasterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `master` | `Invitation` | no |

## GetMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |
| `administratorAccountId` | `string` | no |
| `arn` | `string` | no |
| `email` | `string` | no |
| `invitedAt` | `timestamp` | no |
| `masterAccountId` | `string` | no |
| `relationshipStatus` | `string` | no |
| `tags` | `Map<string>` | no |
| `updatedAt` | `timestamp` | no |

## GetResourceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileUpdatedAt` | `timestamp` | no |
| `sensitivityScore` | `integer` | no |
| `sensitivityScoreOverridden` | `boolean` | no |
| `statistics` | `ResourceStatistics` | no |

## GetRevealConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `RevealConfiguration` | no |
| `retrievalConfiguration` | `RetrievalConfiguration` | no |

## GetSensitiveDataOccurrences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `error` | `string` | no |
| `sensitiveDataOccurrences` | `Map<List<DetectedDataDetails>>` | no |
| `status` | `string` | no |

## GetSensitiveDataOccurrencesAvailability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `code` | `string` | no |
| `reasons` | `List<string>` | no |

## GetSensitivityInspectionTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `excludes` | `SensitivityInspectionTemplateExcludes` | no |
| `includes` | `SensitivityInspectionTemplateIncludes` | no |
| `name` | `string` | no |
| `sensitivityInspectionTemplateId` | `string` | no |

## GetUsageStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterBy` | `List<UsageStatisticsFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `UsageStatisticsSortBy` | no |
| `timeRange` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `records` | `List<UsageRecord>` | no |
| `timeRange` | `string` | no |

## GetUsageTotals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `timeRange` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `timeRange` | `string` | no |
| `usageTotals` | `List<UsageTotal>` | no |

## ListAllowLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `allowLists` | `List<AllowListSummary>` | no |
| `nextToken` | `string` | no |

## ListAutomatedDiscoveryAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AutomatedDiscoveryAccount>` | no |
| `nextToken` | `string` | no |

## ListClassificationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterCriteria` | `ListJobsFilterCriteria` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortCriteria` | `ListJobsSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<JobSummary>` | no |
| `nextToken` | `string` | no |

## ListClassificationScopes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `classificationScopes` | `List<ClassificationScopeSummary>` | no |
| `nextToken` | `string` | no |

## ListCustomDataIdentifiers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<CustomDataIdentifierSummary>` | no |
| `nextToken` | `string` | no |

## ListFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingCriteria` | `FindingCriteria` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingIds` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListFindingsFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingsFilterListItems` | `List<FindingsFilterListItem>` | no |
| `nextToken` | `string` | no |

## ListInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invitations` | `List<Invitation>` | no |
| `nextToken` | `string` | no |

## ListManagedDataIdentifiers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ManagedDataIdentifierSummary>` | no |
| `nextToken` | `string` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `onlyAssociated` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<Member>` | no |
| `nextToken` | `string` | no |

## ListOrganizationAdminAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccounts` | `List<AdminAccount>` | no |
| `nextToken` | `string` | no |

## ListResourceProfileArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifacts` | `List<ResourceProfileArtifact>` | no |
| `nextToken` | `string` | no |

## ListResourceProfileDetections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detections` | `List<Detection>` | no |
| `nextToken` | `string` | no |

## ListSensitivityInspectionTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sensitivityInspectionTemplates` | `List<SensitivityInspectionTemplatesEntry>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutClassificationExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `ClassificationExportConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `ClassificationExportConfiguration` | no |

## PutFindingsPublicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `securityHubConfiguration` | `SecurityHubConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketCriteria` | `SearchResourcesBucketCriteria` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortCriteria` | `SearchResourcesSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `matchingResources` | `List<MatchingResource>` | no |
| `nextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestCustomDataIdentifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ignoreWords` | `List<string>` | no |
| `keywords` | `List<string>` | no |
| `maximumMatchDistance` | `integer` | no |
| `regex` | `string` | yes |
| `sampleText` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `matchCount` | `integer` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAllowList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `criteria` | `AllowListCriteria` | yes |
| `description` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |

## UpdateAutomatedDiscoveryConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnableOrganizationMembers` | `string` | no |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateClassificationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `jobStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateClassificationScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `s3` | `S3ClassificationScopeUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFindingsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | no |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `findingCriteria` | `FindingCriteria` | no |
| `id` | `string` | yes |
| `name` | `string` | no |
| `position` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |

## UpdateMacieSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingPublishingFrequency` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMemberSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnable` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `sensitivityScoreOverride` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourceProfileDetections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `suppressDataIdentifiers` | `List<SuppressDataIdentifier>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRevealConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `RevealConfiguration` | yes |
| `retrievalConfiguration` | `UpdateRetrievalConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `RevealConfiguration` | no |
| `retrievalConfiguration` | `RetrievalConfiguration` | no |

## UpdateSensitivityInspectionTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `excludes` | `SensitivityInspectionTemplateExcludes` | no |
| `id` | `string` | yes |
| `includes` | `SensitivityInspectionTemplateIncludes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


