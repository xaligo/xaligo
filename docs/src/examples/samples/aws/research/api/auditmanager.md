# AWS Audit Manager

API version: 2017-07-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/auditmanager/2017-07-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAssessmentReportEvidenceFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchAssociateAssessmentReportEvidence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |
| `evidenceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidenceIds` | `List<string>` | no |
| `errors` | `List<AssessmentReportEvidenceError>` | no |

## BatchCreateDelegationByAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createDelegationRequests` | `List<CreateDelegationRequest>` | yes |
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegations` | `List<Delegation>` | no |
| `errors` | `List<BatchCreateDelegationByAssessmentError>` | no |

## BatchDeleteDelegationByAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegationIds` | `List<string>` | yes |
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteDelegationByAssessmentError>` | no |

## BatchDisassociateAssessmentReportEvidence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |
| `evidenceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidenceIds` | `List<string>` | no |
| `errors` | `List<AssessmentReportEvidenceError>` | no |

## BatchImportEvidenceToAssessmentControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `controlId` | `string` | yes |
| `manualEvidence` | `List<ManualEvidence>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchImportEvidenceToAssessmentControlError>` | no |

## CreateAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `assessmentReportsDestination` | `AssessmentReportsDestination` | yes |
| `scope` | `Scope` | yes |
| `roles` | `List<Role>` | yes |
| `frameworkId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessment` | `Assessment` | no |

## CreateAssessmentFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `complianceType` | `string` | no |
| `controlSets` | `List<CreateAssessmentFrameworkControlSet>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `framework` | `Framework` | no |

## CreateAssessmentReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `assessmentId` | `string` | yes |
| `queryStatement` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentReport` | `AssessmentReport` | no |

## CreateControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `testingInformation` | `string` | no |
| `actionPlanTitle` | `string` | no |
| `actionPlanInstructions` | `string` | no |
| `controlMappingSources` | `List<CreateControlMappingSource>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `control` | `Control` | no |

## DeleteAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssessmentFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frameworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssessmentFrameworkShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | yes |
| `requestType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssessmentReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `assessmentReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## DeregisterOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateAssessmentReportEvidenceFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## GetAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessment` | `Assessment` | no |
| `userRole` | `Role` | no |

## GetAssessmentFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frameworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `framework` | `Framework` | no |

## GetAssessmentReportUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentReportId` | `string` | yes |
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `preSignedUrl` | `URL` | no |

## GetChangeLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | no |
| `controlId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changeLogs` | `List<ChangeLog>` | no |
| `nextToken` | `string` | no |

## GetControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `control` | `Control` | no |

## GetDelegations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegations` | `List<DelegationMetadata>` | no |
| `nextToken` | `string` | no |

## GetEvidence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |
| `evidenceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidence` | `Evidence` | no |

## GetEvidenceByEvidenceFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidence` | `List<Evidence>` | no |
| `nextToken` | `string` | no |

## GetEvidenceFileUploadUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidenceFileName` | `string` | no |
| `uploadUrl` | `string` | no |

## GetEvidenceFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `evidenceFolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidenceFolder` | `AssessmentEvidenceFolder` | no |

## GetEvidenceFoldersByAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidenceFolders` | `List<AssessmentEvidenceFolder>` | no |
| `nextToken` | `string` | no |

## GetEvidenceFoldersByAssessmentControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `controlId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evidenceFolders` | `List<AssessmentEvidenceFolder>` | no |
| `nextToken` | `string` | no |

## GetInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `insights` | `Insights` | no |

## GetInsightsByAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `insights` | `InsightsByAssessment` | no |

## GetOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccountId` | `string` | no |
| `organizationId` | `string` | no |

## GetServicesInScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceMetadata` | `List<ServiceMetadata>` | no |

## GetSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attribute` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `Settings` | no |

## ListAssessmentControlInsightsByControlDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlDomainId` | `string` | yes |
| `assessmentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlInsightsByAssessment` | `List<ControlInsightsMetadataByAssessmentItem>` | no |
| `nextToken` | `string` | no |

## ListAssessmentFrameworkShareRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentFrameworkShareRequests` | `List<AssessmentFrameworkShareRequest>` | no |
| `nextToken` | `string` | no |

## ListAssessmentFrameworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frameworkType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frameworkMetadataList` | `List<AssessmentFrameworkMetadata>` | no |
| `nextToken` | `string` | no |

## ListAssessmentReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentReports` | `List<AssessmentReportMetadata>` | no |
| `nextToken` | `string` | no |

## ListAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentMetadata` | `List<AssessmentMetadataItem>` | no |
| `nextToken` | `string` | no |

## ListControlDomainInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlDomainInsights` | `List<ControlDomainInsights>` | no |
| `nextToken` | `string` | no |

## ListControlDomainInsightsByAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlDomainInsights` | `List<ControlDomainInsights>` | no |
| `nextToken` | `string` | no |

## ListControlInsightsByControlDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlDomainId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlInsightsMetadata` | `List<ControlInsightsMetadataItem>` | no |
| `nextToken` | `string` | no |

## ListControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `controlCatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlMetadataList` | `List<ControlMetadata>` | no |
| `nextToken` | `string` | no |

## ListKeywordsForDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `source` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keywords` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notifications` | `List<Notification>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RegisterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKey` | `string` | no |
| `delegatedAdminAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## RegisterOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminAccountId` | `string` | no |
| `organizationId` | `string` | no |

## StartAssessmentFrameworkShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frameworkId` | `string` | yes |
| `destinationAccount` | `string` | yes |
| `destinationRegion` | `string` | yes |
| `comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentFrameworkShareRequest` | `AssessmentFrameworkShareRequest` | no |

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


## UpdateAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `assessmentName` | `string` | no |
| `assessmentDescription` | `string` | no |
| `scope` | `Scope` | yes |
| `assessmentReportsDestination` | `AssessmentReportsDestination` | no |
| `roles` | `List<Role>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessment` | `Assessment` | no |

## UpdateAssessmentControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `controlId` | `string` | yes |
| `controlStatus` | `string` | no |
| `commentBody` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `control` | `AssessmentControl` | no |

## UpdateAssessmentControlSetStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `controlSetId` | `string` | yes |
| `status` | `string` | yes |
| `comment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlSet` | `AssessmentControlSet` | no |

## UpdateAssessmentFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frameworkId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `complianceType` | `string` | no |
| `controlSets` | `List<UpdateAssessmentFrameworkControlSet>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `framework` | `Framework` | no |

## UpdateAssessmentFrameworkShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | yes |
| `requestType` | `string` | yes |
| `action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentFrameworkShareRequest` | `AssessmentFrameworkShareRequest` | no |

## UpdateAssessmentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessment` | `Assessment` | no |

## UpdateControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `testingInformation` | `string` | no |
| `actionPlanTitle` | `string` | no |
| `actionPlanInstructions` | `string` | no |
| `controlMappingSources` | `List<ControlMappingSource>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `control` | `Control` | no |

## UpdateSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snsTopic` | `string` | no |
| `defaultAssessmentReportsDestination` | `AssessmentReportsDestination` | no |
| `defaultProcessOwners` | `List<Role>` | no |
| `kmsKey` | `string` | no |
| `evidenceFinderEnabled` | `boolean` | no |
| `deregistrationPolicy` | `DeregistrationPolicy` | no |
| `defaultExportDestination` | `DefaultExportDestination` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `Settings` | no |

## ValidateAssessmentReportIntegrity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3RelativePath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signatureValid` | `boolean` | no |
| `signatureAlgorithm` | `string` | no |
| `signatureDateTime` | `string` | no |
| `signatureKeyId` | `string` | no |
| `validationErrors` | `List<string>` | no |

