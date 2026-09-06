# Partner Central Selling API

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/partnercentral-selling/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptEngagementInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssignOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `Assignee` | `AssigneeContact` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `OpportunityIdentifier` | `string` | yes |
| `RelatedEntityType` | `string` | yes |
| `RelatedEntityIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Title` | `string` | no |
| `Description` | `string` | no |
| `Contexts` | `List<EngagementContextDetails>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `ModifiedAt` | `timestamp` | no |

## CreateEngagementContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `EngagementIdentifier` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Type` | `string` | yes |
| `Payload` | `EngagementContextPayload` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementId` | `string` | no |
| `EngagementArn` | `string` | no |
| `EngagementLastModifiedAt` | `timestamp` | no |
| `ContextId` | `string` | no |

## CreateEngagementInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `EngagementIdentifier` | `string` | yes |
| `Invitation` | `Invitation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |

## CreateOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `PrimaryNeedsFromAws` | `List<string>` | no |
| `NationalSecurity` | `string` | no |
| `PartnerOpportunityIdentifier` | `string` | no |
| `Customer` | `Customer` | no |
| `Project` | `Project` | no |
| `OpportunityType` | `string` | no |
| `Marketing` | `Marketing` | no |
| `SoftwareRevenue` | `SoftwareRevenue` | no |
| `ClientToken` | `string` | yes |
| `LifeCycle` | `LifeCycle` | no |
| `Origin` | `string` | no |
| `OpportunityTeam` | `List<Contact>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `PartnerOpportunityIdentifier` | `string` | no |
| `LastModifiedDate` | `timestamp` | no |

## CreateResourceSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `EngagementIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ResourceIdentifier` | `string` | yes |
| `ResourceSnapshotTemplateIdentifier` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Revision` | `integer` | no |

## CreateResourceSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `EngagementIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ResourceIdentifier` | `string` | yes |
| `ResourceSnapshotTemplateIdentifier` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## DeleteResourceSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `OpportunityIdentifier` | `string` | yes |
| `RelatedEntityType` | `string` | yes |
| `RelatedEntityIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAwsOpportunitySummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `RelatedOpportunityIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelatedOpportunityId` | `string` | no |
| `Origin` | `string` | no |
| `InvolvementType` | `string` | no |
| `Visibility` | `string` | no |
| `LifeCycle` | `AwsOpportunityLifeCycle` | no |
| `OpportunityTeam` | `List<AwsTeamMember>` | no |
| `Insights` | `AwsOpportunityInsights` | no |
| `InvolvementTypeChangeReason` | `string` | no |
| `RelatedEntityIds` | `AwsOpportunityRelatedEntities` | no |
| `Customer` | `AwsOpportunityCustomer` | no |
| `Project` | `AwsOpportunityProject` | no |
| `CosellMotion` | `string` | no |
| `SoftwareRevenue` | `AwsSoftwareRevenue` | no |
| `Catalog` | `string` | yes |

## GetEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Title` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `MemberCount` | `integer` | no |
| `ModifiedAt` | `timestamp` | no |
| `ModifiedBy` | `string` | no |
| `Contexts` | `List<EngagementContextDetails>` | no |

## GetEngagementInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `PayloadType` | `string` | no |
| `Id` | `string` | yes |
| `EngagementId` | `string` | no |
| `EngagementTitle` | `string` | no |
| `Status` | `string` | no |
| `InvitationDate` | `timestamp` | no |
| `ExpirationDate` | `timestamp` | no |
| `SenderAwsAccountId` | `string` | no |
| `SenderCompanyName` | `string` | no |
| `Receiver` | `Receiver` | no |
| `Catalog` | `string` | yes |
| `RejectionReason` | `string` | no |
| `Payload` | `Payload` | no |
| `InvitationMessage` | `string` | no |
| `EngagementDescription` | `string` | no |
| `ExistingMembers` | `List<EngagementMemberSummary>` | no |
| `EnrichmentContext` | `EnrichmentContext` | no |

## GetOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `PrimaryNeedsFromAws` | `List<string>` | no |
| `NationalSecurity` | `string` | no |
| `PartnerOpportunityIdentifier` | `string` | no |
| `Customer` | `Customer` | no |
| `Project` | `Project` | no |
| `OpportunityType` | `string` | no |
| `Marketing` | `Marketing` | no |
| `SoftwareRevenue` | `SoftwareRevenue` | no |
| `Id` | `string` | yes |
| `Arn` | `string` | no |
| `LastModifiedDate` | `timestamp` | yes |
| `CreatedDate` | `timestamp` | yes |
| `RelatedEntityIdentifiers` | `RelatedEntityIdentifiers` | yes |
| `LifeCycle` | `LifeCycle` | no |
| `OpportunityTeam` | `List<Contact>` | no |

## GetProspectingFromEngagementTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `TaskIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |
| `TaskArn` | `string` | yes |
| `TaskName` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | no |
| `Engagements` | `List<EngagementProspectingResult>` | yes |

## GetResourceSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `EngagementIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ResourceIdentifier` | `string` | yes |
| `ResourceSnapshotTemplateIdentifier` | `string` | yes |
| `Revision` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | no |
| `CreatedBy` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `EngagementId` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceId` | `string` | no |
| `ResourceSnapshotTemplateName` | `string` | no |
| `Revision` | `integer` | no |
| `Payload` | `ResourceSnapshotPayload` | no |
| `TargetMemberAccounts` | `List<string>` | no |

## GetResourceSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `EngagementId` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceId` | `string` | no |
| `ResourceArn` | `string` | no |
| `ResourceSnapshotTemplateName` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Status` | `string` | no |
| `LastSuccessfulExecutionDate` | `timestamp` | no |
| `LastFailure` | `string` | no |

## GetSellingSystemSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobRoleArn` | `string` | no |

## ListEngagementByAcceptingInvitationTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `ListTasksSortBase` | no |
| `Catalog` | `string` | yes |
| `TaskStatus` | `List<string>` | no |
| `OpportunityIdentifier` | `List<string>` | no |
| `EngagementInvitationIdentifier` | `List<string>` | no |
| `TaskIdentifier` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskSummaries` | `List<ListEngagementByAcceptingInvitationTaskSummary>` | no |
| `NextToken` | `string` | no |

## ListEngagementFromOpportunityTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `ListTasksSortBase` | no |
| `Catalog` | `string` | yes |
| `TaskStatus` | `List<string>` | no |
| `TaskIdentifier` | `List<string>` | no |
| `OpportunityIdentifier` | `List<string>` | no |
| `EngagementIdentifier` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskSummaries` | `List<ListEngagementFromOpportunityTaskSummary>` | no |
| `NextToken` | `string` | no |

## ListEngagementInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `OpportunityEngagementInvitationSort` | no |
| `PayloadType` | `List<string>` | no |
| `ParticipantType` | `string` | yes |
| `Status` | `List<string>` | no |
| `EngagementIdentifier` | `List<string>` | no |
| `SenderAwsAccountId` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementInvitationSummaries` | `List<EngagementInvitationSummary>` | no |
| `NextToken` | `string` | no |

## ListEngagementMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementMemberList` | `List<EngagementMember>` | yes |
| `NextToken` | `string` | no |

## ListEngagementResourceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `EngagementIdentifier` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceIdentifier` | `string` | no |
| `CreatedBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementResourceAssociationSummaries` | `List<EngagementResourceAssociationSummary>` | yes |
| `NextToken` | `string` | no |

## ListEngagements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `CreatedBy` | `List<string>` | no |
| `ExcludeCreatedBy` | `List<string>` | no |
| `ContextTypes` | `List<string>` | no |
| `ExcludeContextTypes` | `List<string>` | no |
| `Sort` | `EngagementSort` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `EngagementIdentifier` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementSummaryList` | `List<EngagementSummary>` | yes |
| `NextToken` | `string` | no |

## ListOpportunities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `OpportunitySort` | no |
| `LastModifiedDate` | `LastModifiedDate` | no |
| `Identifier` | `List<string>` | no |
| `LifeCycleStage` | `List<string>` | no |
| `LifeCycleReviewStatus` | `List<string>` | no |
| `CustomerCompanyName` | `List<string>` | no |
| `CreatedDate` | `CreatedDateFilter` | no |
| `TargetCloseDate` | `TargetCloseDateFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpportunitySummaries` | `List<OpportunitySummary>` | yes |
| `NextToken` | `string` | no |

## ListOpportunityFromEngagementTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `ListTasksSortBase` | no |
| `Catalog` | `string` | yes |
| `TaskStatus` | `List<string>` | no |
| `TaskIdentifier` | `List<string>` | no |
| `OpportunityIdentifier` | `List<string>` | no |
| `EngagementIdentifier` | `List<string>` | no |
| `ContextIdentifier` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskSummaries` | `List<ListOpportunityFromEngagementTaskSummary>` | no |
| `NextToken` | `string` | no |

## ListProspectingFromEngagementTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TaskIdentifier` | `List<string>` | no |
| `TaskName` | `List<string>` | no |
| `StartAfter` | `timestamp` | no |
| `StartBefore` | `timestamp` | no |
| `Sort` | `ProspectingFromEngagementTaskSort` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TaskSummaries` | `List<ProspectingTaskSummary>` | yes |

## ListResourceSnapshotJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `EngagementIdentifier` | `string` | no |
| `Status` | `string` | no |
| `Sort` | `SortObject` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSnapshotJobSummaries` | `List<ResourceSnapshotJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListResourceSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `EngagementIdentifier` | `string` | yes |
| `ResourceType` | `string` | no |
| `ResourceIdentifier` | `string` | no |
| `ResourceSnapshotTemplateIdentifier` | `string` | no |
| `CreatedBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSnapshotSummaries` | `List<ResourceSnapshotSummary>` | yes |
| `NextToken` | `string` | no |

## ListSolutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `SolutionSort` | no |
| `Status` | `List<string>` | no |
| `Identifier` | `List<string>` | no |
| `Category` | `List<string>` | no |
| `AwsMarketplaceSolutionArn` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SolutionSummaries` | `List<SolutionBase>` | yes |
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

## PutSellingSystemSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobRoleIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobRoleArn` | `string` | no |

## RejectEngagementInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `RejectionReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartEngagementByAcceptingInvitationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Identifier` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | no |
| `TaskArn` | `string` | no |
| `StartTime` | `timestamp` | no |
| `TaskStatus` | `string` | no |
| `Message` | `string` | no |
| `ReasonCode` | `string` | no |
| `OpportunityId` | `string` | no |
| `ResourceSnapshotJobId` | `string` | no |
| `EngagementInvitationId` | `string` | no |

## StartEngagementFromOpportunityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Identifier` | `string` | yes |
| `AwsSubmission` | `AwsSubmission` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | no |
| `TaskArn` | `string` | no |
| `StartTime` | `timestamp` | no |
| `TaskStatus` | `string` | no |
| `Message` | `string` | no |
| `ReasonCode` | `string` | no |
| `OpportunityId` | `string` | no |
| `ResourceSnapshotJobId` | `string` | no |
| `EngagementId` | `string` | no |
| `EngagementInvitationId` | `string` | no |

## StartOpportunityFromEngagementTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Identifier` | `string` | yes |
| `ContextIdentifier` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | no |
| `TaskArn` | `string` | no |
| `StartTime` | `timestamp` | no |
| `TaskStatus` | `string` | no |
| `Message` | `string` | no |
| `ReasonCode` | `string` | no |
| `OpportunityId` | `string` | no |
| `ResourceSnapshotJobId` | `string` | no |
| `EngagementId` | `string` | no |
| `ContextId` | `string` | no |

## StartProspectingFromEngagementTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifiers` | `List<string>` | yes |
| `TaskName` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifiers` | `List<string>` | yes |
| `TaskName` | `string` | yes |
| `Message` | `string` | no |
| `ReasonCode` | `string` | no |
| `StartTime` | `timestamp` | yes |
| `TaskId` | `string` | no |
| `TaskArn` | `string` | no |
| `TaskStatus` | `string` | yes |

## StartResourceSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopResourceSnapshotJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ResourceSnapshotJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubmitOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `InvolvementType` | `string` | yes |
| `Visibility` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateEngagementContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `EngagementIdentifier` | `string` | yes |
| `ContextIdentifier` | `string` | yes |
| `EngagementLastModifiedAt` | `timestamp` | yes |
| `Type` | `string` | yes |
| `Payload` | `UpdateEngagementContextPayload` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementId` | `string` | yes |
| `EngagementArn` | `string` | yes |
| `EngagementLastModifiedAt` | `timestamp` | yes |
| `ContextId` | `string` | yes |

## UpdateOpportunity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `PrimaryNeedsFromAws` | `List<string>` | no |
| `NationalSecurity` | `string` | no |
| `PartnerOpportunityIdentifier` | `string` | no |
| `Customer` | `Customer` | no |
| `Project` | `Project` | no |
| `OpportunityType` | `string` | no |
| `Marketing` | `Marketing` | no |
| `SoftwareRevenue` | `SoftwareRevenue` | no |
| `LastModifiedDate` | `timestamp` | yes |
| `Identifier` | `string` | yes |
| `LifeCycle` | `LifeCycle` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `LastModifiedDate` | `timestamp` | yes |

