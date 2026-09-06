# AWS Support

API version: 2013-04-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/support/2013-04-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddAttachmentsToSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentSetId` | `string` | no |
| `attachments` | `List<Attachment>` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentSetId` | `string` | no |
| `expiryTime` | `string` | no |

## AddCommunicationToCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | no |
| `communicationBody` | `string` | yes |
| `ccEmailAddresses` | `List<string>` | no |
| `attachmentSetId` | `string` | no |
| `uploadIds` | `List<string>` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `boolean` | no |

## CompleteAttachmentUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |
| `completedUploads` | `List<CompletedUpload>` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadStatus` | `string` | yes |

## CreateCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subject` | `string` | yes |
| `serviceCode` | `string` | no |
| `severityCode` | `string` | no |
| `categoryCode` | `string` | no |
| `communicationBody` | `string` | yes |
| `ccEmailAddresses` | `List<string>` | no |
| `language` | `string` | no |
| `issueType` | `string` | no |
| `attachmentSetId` | `string` | no |
| `uploadIds` | `List<string>` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | no |

## DescribeAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentId` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachment` | `Attachment` | no |

## DescribeAttachmentUploadStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadStatus` | `string` | yes |
| `fileName` | `string` | yes |
| `uploadProgress` | `UploadProgress` | no |

## DescribeCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseIdList` | `List<string>` | no |
| `displayId` | `string` | no |
| `afterTime` | `string` | no |
| `beforeTime` | `string` | no |
| `includeResolvedCases` | `boolean` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `language` | `string` | no |
| `includeCommunications` | `boolean` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cases` | `List<CaseDetails>` | no |
| `nextToken` | `string` | no |

## DescribeCommunications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `beforeTime` | `string` | no |
| `afterTime` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `communications` | `List<Communication>` | no |
| `nextToken` | `string` | no |

## DescribeCreateCaseOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `issueType` | `string` | yes |
| `serviceCode` | `string` | yes |
| `language` | `string` | yes |
| `categoryCode` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `languageAvailability` | `string` | no |
| `communicationTypes` | `List<CommunicationTypeOptions>` | no |

## DescribeServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceCodeList` | `List<string>` | no |
| `language` | `string` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `services` | `List<Service>` | no |

## DescribeSeverityLevels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `language` | `string` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `severityLevels` | `List<SeverityLevel>` | no |

## DescribeSupportedLanguages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `issueType` | `string` | yes |
| `serviceCode` | `string` | yes |
| `categoryCode` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `supportedLanguages` | `List<SupportedLanguage>` | no |

## DescribeTrustedAdvisorCheckRefreshStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statuses` | `List<TrustedAdvisorCheckRefreshStatus>` | yes |

## DescribeTrustedAdvisorCheckResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkId` | `string` | yes |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `TrustedAdvisorCheckResult` | no |

## DescribeTrustedAdvisorCheckSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<TrustedAdvisorCheckSummary>` | yes |

## DescribeTrustedAdvisorChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `language` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checks` | `List<TrustedAdvisorCheckDescription>` | yes |

## GetAttachmentDownloadLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentId` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileName` | `string` | yes |
| `downloadUrl` | `DownloadUrl` | yes |

## GetAttachmentUploadLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileName` | `string` | yes |
| `fileSizeBytes` | `long` | no |
| `uploadId` | `string` | no |
| `uploadRange` | `UploadRange` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |
| `partSizeBytes` | `long` | yes |
| `totalParts` | `integer` | yes |
| `nextIndex` | `integer` | no |
| `uploadUrls` | `List<UploadUrl>` | yes |

## RefreshTrustedAdvisorCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `TrustedAdvisorCheckRefreshStatus` | yes |

## ResolveCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `initialCaseStatus` | `string` | no |
| `finalCaseStatus` | `string` | no |

