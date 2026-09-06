# AWS Well-Architected Tool

API version: 2020-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/wellarchitected/2020-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateLenses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAliases` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `ProfileArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAgentContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `profileArn` | `string` | yes |
| `title` | `string` | yes |
| `contextType` | `string` | yes |
| `content` | `ContextContent` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `context` | `ContextSummary` | yes |

## CreateAgentGoal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `profileArn` | `string` | yes |
| `pillars` | `List<string>` | yes |
| `title` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `goal` | `GoalSummary` | yes |

## CreateAgentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `businessOverview` | `string` | no |
| `pillars` | `List<string>` | yes |
| `deletionProtection` | `boolean` | no |
| `executionRoleArn` | `string` | yes |
| `aggregationConfiguration` | `List<AggregationConfiguration>` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `businessOverview` | `string` | no |
| `pillars` | `List<string>` | yes |
| `deletionProtection` | `boolean` | no |
| `executionRoleArn` | `string` | yes |
| `aggregationConfiguration` | `List<AggregationConfiguration>` | yes |
| `arn` | `string` | yes |
| `eligibleForScheduledGeneration` | `boolean` | no |
| `eligibleForArchitectureGeneration` | `boolean` | no |
| `fieldErrors` | `Map<string>` | no |
| `tags` | `List<Tag>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `timestamp` | no |

## CreateLensShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `SharedWith` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | no |

## CreateLensVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `LensVersion` | `string` | yes |
| `IsMajorVersion` | `boolean` | no |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensArn` | `string` | no |
| `LensVersion` | `string` | no |

## CreateMilestone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `MilestoneName` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |

## CreateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileName` | `string` | yes |
| `ProfileDescription` | `string` | yes |
| `ProfileQuestions` | `List<ProfileQuestionUpdate>` | yes |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileArn` | `string` | no |
| `ProfileVersion` | `string` | no |

## CreateProfileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileArn` | `string` | yes |
| `SharedWith` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | no |
| `ProfileArn` | `string` | no |

## CreateReviewTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Description` | `string` | yes |
| `Lenses` | `List<string>` | yes |
| `Notes` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |

## CreateTemplateShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `SharedWith` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `ShareId` | `string` | no |

## CreateWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadName` | `string` | yes |
| `Description` | `string` | yes |
| `Environment` | `string` | yes |
| `AccountIds` | `List<string>` | no |
| `AwsRegions` | `List<string>` | no |
| `NonAwsRegions` | `List<string>` | no |
| `PillarPriorities` | `List<string>` | no |
| `ArchitecturalDesign` | `string` | no |
| `ReviewOwner` | `string` | no |
| `IndustryType` | `string` | no |
| `Industry` | `string` | no |
| `Lenses` | `List<string>` | yes |
| `Notes` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `DiscoveryConfig` | `WorkloadDiscoveryConfig` | no |
| `Applications` | `List<string>` | no |
| `ProfileArns` | `List<string>` | no |
| `ReviewTemplateArns` | `List<string>` | no |
| `JiraConfiguration` | `WorkloadJiraConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `WorkloadArn` | `string` | no |

## CreateWorkloadShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `SharedWith` | `string` | yes |
| `PermissionType` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `ShareId` | `string` | no |

## DeleteAgentContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAgentGoal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAgentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `ClientRequestToken` | `string` | yes |
| `LensStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLensShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileArn` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProfileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | yes |
| `ProfileArn` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReviewTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTemplateShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | yes |
| `TemplateArn` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkloadShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | yes |
| `WorkloadId` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateLenses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAliases` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `ProfileArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportLens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `LensVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensJSON` | `string` | no |

## GetAgentContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `context` | `ContextSummary` | yes |

## GetAgentGoal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `goal` | `GoalSummary` | yes |

## GetAgentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `businessOverview` | `string` | no |
| `pillars` | `List<string>` | yes |
| `deletionProtection` | `boolean` | no |
| `executionRoleArn` | `string` | yes |
| `aggregationConfiguration` | `List<AggregationConfiguration>` | yes |
| `arn` | `string` | yes |
| `eligibleForScheduledGeneration` | `boolean` | no |
| `eligibleForArchitectureGeneration` | `boolean` | no |
| `fieldErrors` | `Map<string>` | no |
| `tags` | `List<Tag>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `timestamp` | no |

## GetAgentRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationArn` | `string` | yes |
| `remediationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationArn` | `string` | yes |
| `profileArn` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | yes |
| `type` | `string` | yes |
| `pillar` | `string` | yes |
| `priority` | `string` | yes |
| `effort` | `string` | yes |
| `status` | `string` | yes |
| `state` | `string` | yes |
| `updateReason` | `string` | no |
| `impact` | `string` | yes |
| `roi` | `Roi` | yes |
| `numberOfResources` | `integer` | no |
| `awsServices` | `List<string>` | no |
| `businessUnits` | `List<string>` | no |
| `applications` | `List<string>` | no |
| `impactDetails` | `List<string>` | yes |
| `insights` | `List<Insight>` | yes |
| `highlights` | `List<string>` | yes |
| `remediationSummary` | `RemediationSummary` | yes |
| `crossPillarBenefits` | `List<CrossPillarBenefit>` | no |
| `tradeOffs` | `List<TradeOff>` | no |
| `sources` | `List<string>` | no |
| `goals` | `List<RecommendationGoal>` | no |
| `tags` | `List<Tag>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `timestamp` | no |
| `remediations` | `List<AgentRecommendationRemediation>` | no |

## GetAgentRecommendationGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `generationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `profileArn` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | yes |
| `estimatedCompletionTime` | `timestamp` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `timestamp` | no |
| `additionalContext` | `Document` | no |
| `scope` | `Scope` | no |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `progress` | `Progress` | no |
| `errorDetails` | `ErrorDetails` | no |

## GetAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `QuestionId` | `string` | yes |
| `MilestoneNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `LensAlias` | `string` | no |
| `LensArn` | `string` | no |
| `Answer` | `Answer` | no |

## GetConsolidatedReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Format` | `string` | yes |
| `IncludeSharedResources` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metrics` | `List<ConsolidatedReportMetric>` | no |
| `NextToken` | `string` | no |
| `Base64String` | `string` | no |

## GetGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationSharingStatus` | `string` | no |
| `DiscoveryIntegrationStatus` | `string` | no |
| `JiraConfiguration` | `AccountJiraConfigurationOutput` | no |

## GetLens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `LensVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Lens` | `Lens` | no |

## GetLensReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `MilestoneNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `LensReview` | `LensReview` | no |

## GetLensReviewReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `MilestoneNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `LensReviewReport` | `LensReviewReport` | no |

## GetLensVersionDifference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `BaseLensVersion` | `string` | no |
| `TargetLensVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | no |
| `LensArn` | `string` | no |
| `BaseLensVersion` | `string` | no |
| `TargetLensVersion` | `string` | no |
| `LatestLensVersion` | `string` | no |
| `VersionDifferences` | `VersionDifferences` | no |

## GetMilestone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `MilestoneNumber` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `Milestone` | `Milestone` | no |

## GetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileArn` | `string` | yes |
| `ProfileVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profile` | `Profile` | no |

## GetProfileTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileTemplate` | `ProfileTemplate` | no |

## GetReviewTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReviewTemplate` | `ReviewTemplate` | no |

## GetReviewTemplateAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `LensAlias` | `string` | yes |
| `QuestionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `LensAlias` | `string` | no |
| `Answer` | `ReviewTemplateAnswer` | no |

## GetReviewTemplateLensReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `LensAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `LensReview` | `ReviewTemplateLensReview` | no |

## GetWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workload` | `Workload` | no |

## ImportLens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | no |
| `JSONString` | `string` | yes |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensArn` | `string` | no |
| `Status` | `string` | no |

## ListAgentContexts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ContextSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentGoals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<GoalSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AgentProfileSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentRecommendationGenerations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `recommendationType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AgentRecommendationGenerationSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentRecommendationItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationArn` | `string` | yes |
| `type` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AgentRecommendationItemSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `state` | `string` | no |
| `pillar` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AgentRecommendationSummary>` | yes |
| `nextToken` | `string` | no |

## ListAnswers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `PillarId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `QuestionPriority` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `LensAlias` | `string` | no |
| `LensArn` | `string` | no |
| `AnswerSummaries` | `List<AnswerSummary>` | no |
| `NextToken` | `string` | no |

## ListCheckDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LensArn` | `string` | yes |
| `PillarId` | `string` | yes |
| `QuestionId` | `string` | yes |
| `ChoiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CheckDetails` | `List<CheckDetail>` | no |
| `NextToken` | `string` | no |

## ListCheckSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LensArn` | `string` | yes |
| `PillarId` | `string` | yes |
| `QuestionId` | `string` | yes |
| `ChoiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CheckSummaries` | `List<CheckSummary>` | no |
| `NextToken` | `string` | no |

## ListLensReviewImprovements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `PillarId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `QuestionPriority` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `LensAlias` | `string` | no |
| `LensArn` | `string` | no |
| `ImprovementSummaries` | `List<ImprovementSummary>` | no |
| `NextToken` | `string` | no |

## ListLensReviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `MilestoneNumber` | `integer` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneNumber` | `integer` | no |
| `LensReviewSummaries` | `List<LensReviewSummary>` | no |
| `NextToken` | `string` | no |

## ListLensShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensAlias` | `string` | yes |
| `SharedWithPrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensShareSummaries` | `List<LensShareSummary>` | no |
| `NextToken` | `string` | no |

## ListLenses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LensType` | `string` | no |
| `LensStatus` | `string` | no |
| `LensName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LensSummaries` | `List<LensSummary>` | no |
| `NextToken` | `string` | no |

## ListMilestones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `MilestoneSummaries` | `List<MilestoneSummary>` | no |
| `NextToken` | `string` | no |

## ListNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationSummaries` | `List<NotificationSummary>` | no |
| `NextToken` | `string` | no |

## ListProfileNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationSummaries` | `List<ProfileNotificationSummary>` | no |
| `NextToken` | `string` | no |

## ListProfileShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileArn` | `string` | yes |
| `SharedWithPrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileShareSummaries` | `List<ProfileShareSummary>` | no |
| `NextToken` | `string` | no |

## ListProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileNamePrefix` | `string` | no |
| `ProfileOwnerType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileSummaries` | `List<ProfileSummary>` | no |
| `NextToken` | `string` | no |

## ListReviewTemplateAnswers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `LensAlias` | `string` | yes |
| `PillarId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `LensAlias` | `string` | no |
| `AnswerSummaries` | `List<ReviewTemplateAnswerSummary>` | no |
| `NextToken` | `string` | no |

## ListReviewTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReviewTemplates` | `List<ReviewTemplateSummary>` | no |
| `NextToken` | `string` | no |

## ListShareInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadNamePrefix` | `string` | no |
| `LensNamePrefix` | `string` | no |
| `ShareResourceType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ProfileNamePrefix` | `string` | no |
| `TemplateNamePrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareInvitationSummaries` | `List<ShareInvitationSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTemplateShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `SharedWithPrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `TemplateShareSummaries` | `List<TemplateShareSummary>` | no |
| `NextToken` | `string` | no |

## ListWorkloadShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `SharedWithPrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `WorkloadShareSummaries` | `List<WorkloadShareSummary>` | no |
| `NextToken` | `string` | no |

## ListWorkloads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadNamePrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadSummaries` | `List<WorkloadSummary>` | no |
| `NextToken` | `string` | no |

## PutAgentRecommendationFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationArn` | `string` | yes |
| `type` | `string` | yes |
| `feedbackCategory` | `string` | no |
| `comments` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAgentRecommendationGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileArn` | `string` | yes |
| `types` | `List<string>` | yes |
| `name` | `string` | no |
| `additionalContext` | `Document` | no |
| `scope` | `Scope` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `profileArn` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | yes |
| `estimatedCompletionTime` | `timestamp` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `timestamp` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAgentContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `profileArn` | `string` | yes |
| `id` | `string` | yes |
| `title` | `string` | no |
| `content` | `ContextContent` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `context` | `ContextSummary` | yes |

## UpdateAgentGoal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `profileArn` | `string` | yes |
| `id` | `string` | yes |
| `pillars` | `List<string>` | no |
| `title` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `goal` | `GoalSummary` | yes |

## UpdateAgentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `profileArn` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `executionRoleArn` | `string` | no |
| `aggregationConfiguration` | `List<AggregationConfiguration>` | no |
| `businessOverview` | `string` | no |
| `pillars` | `List<string>` | no |
| `deletionProtection` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `businessOverview` | `string` | no |
| `pillars` | `List<string>` | yes |
| `deletionProtection` | `boolean` | no |
| `executionRoleArn` | `string` | yes |
| `aggregationConfiguration` | `List<AggregationConfiguration>` | yes |
| `arn` | `string` | yes |
| `eligibleForScheduledGeneration` | `boolean` | no |
| `eligibleForArchitectureGeneration` | `boolean` | no |
| `fieldErrors` | `Map<string>` | no |
| `tags` | `List<Tag>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `timestamp` | no |

## UpdateAgentRecommendationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationArn` | `string` | yes |
| `status` | `string` | yes |
| `updateReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `QuestionId` | `string` | yes |
| `SelectedChoices` | `List<string>` | no |
| `ChoiceUpdates` | `Map<ChoiceUpdate>` | no |
| `Notes` | `string` | no |
| `IsApplicable` | `boolean` | no |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `LensAlias` | `string` | no |
| `LensArn` | `string` | no |
| `Answer` | `Answer` | no |

## UpdateGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationSharingStatus` | `string` | no |
| `DiscoveryIntegrationStatus` | `string` | no |
| `JiraConfiguration` | `AccountJiraConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `ClientRequestToken` | `string` | yes |
| `IntegratingService` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLensReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `LensNotes` | `string` | no |
| `PillarNotes` | `Map<string>` | no |
| `JiraConfiguration` | `JiraSelectedQuestionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `LensReview` | `LensReview` | no |

## UpdateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileArn` | `string` | yes |
| `ProfileDescription` | `string` | no |
| `ProfileQuestions` | `List<ProfileQuestionUpdate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profile` | `Profile` | no |

## UpdateReviewTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `TemplateName` | `string` | no |
| `Description` | `string` | no |
| `Notes` | `string` | no |
| `LensesToAssociate` | `List<string>` | no |
| `LensesToDisassociate` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReviewTemplate` | `ReviewTemplate` | no |

## UpdateReviewTemplateAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `LensAlias` | `string` | yes |
| `QuestionId` | `string` | yes |
| `SelectedChoices` | `List<string>` | no |
| `ChoiceUpdates` | `Map<ChoiceUpdate>` | no |
| `Notes` | `string` | no |
| `IsApplicable` | `boolean` | no |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `LensAlias` | `string` | no |
| `Answer` | `ReviewTemplateAnswer` | no |

## UpdateReviewTemplateLensReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `LensAlias` | `string` | yes |
| `LensNotes` | `string` | no |
| `PillarNotes` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |
| `LensReview` | `ReviewTemplateLensReview` | no |

## UpdateShareInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareInvitationId` | `string` | yes |
| `ShareInvitationAction` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareInvitation` | `ShareInvitation` | no |

## UpdateWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `WorkloadName` | `string` | no |
| `Description` | `string` | no |
| `Environment` | `string` | no |
| `AccountIds` | `List<string>` | no |
| `AwsRegions` | `List<string>` | no |
| `NonAwsRegions` | `List<string>` | no |
| `PillarPriorities` | `List<string>` | no |
| `ArchitecturalDesign` | `string` | no |
| `ReviewOwner` | `string` | no |
| `IsReviewOwnerUpdateAcknowledged` | `boolean` | no |
| `IndustryType` | `string` | no |
| `Industry` | `string` | no |
| `Notes` | `string` | no |
| `ImprovementStatus` | `string` | no |
| `DiscoveryConfig` | `WorkloadDiscoveryConfig` | no |
| `Applications` | `List<string>` | no |
| `JiraConfiguration` | `WorkloadJiraConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workload` | `Workload` | no |

## UpdateWorkloadShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareId` | `string` | yes |
| `WorkloadId` | `string` | yes |
| `PermissionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `WorkloadShare` | `WorkloadShare` | no |

## UpgradeLensReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `LensAlias` | `string` | yes |
| `MilestoneName` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpgradeProfileVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | yes |
| `ProfileArn` | `string` | yes |
| `MilestoneName` | `string` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpgradeReviewTemplateLensReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `LensAlias` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


