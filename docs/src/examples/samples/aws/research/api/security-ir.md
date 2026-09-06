# Security Incident Response

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/security-ir/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetMemberAccountDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |
| `accountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<GetMembershipAccountDetailItem>` | no |
| `errors` | `List<GetMembershipAccountDetailError>` | no |

## CancelMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |

## CloseCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseStatus` | `string` | no |
| `closedDate` | `timestamp` | no |

## CreateCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `resolverType` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | yes |
| `engagementType` | `string` | yes |
| `reportedIncidentStartDate` | `timestamp` | yes |
| `impactedAccounts` | `List<string>` | yes |
| `watchers` | `List<Watcher>` | yes |
| `threatActorIpAddresses` | `List<ThreatActorIp>` | no |
| `impactedServices` | `List<string>` | no |
| `impactedAwsRegions` | `List<ImpactedAwsRegion>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |

## CreateCaseComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `clientToken` | `string` | no |
| `body` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |

## CreateMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `membershipName` | `string` | yes |
| `incidentResponseTeam` | `List<IncidentResponder>` | yes |
| `optInFeatures` | `List<OptInFeature>` | no |
| `tags` | `Map<string>` | no |
| `coverEntireOrganization` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |

## GetCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | no |
| `caseArn` | `string` | no |
| `description` | `string` | no |
| `caseStatus` | `string` | no |
| `engagementType` | `string` | no |
| `reportedIncidentStartDate` | `timestamp` | no |
| `actualIncidentStartDate` | `timestamp` | no |
| `impactedAwsRegions` | `List<ImpactedAwsRegion>` | no |
| `threatActorIpAddresses` | `List<ThreatActorIp>` | no |
| `pendingAction` | `string` | no |
| `impactedAccounts` | `List<string>` | no |
| `watchers` | `List<Watcher>` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `closureCode` | `string` | no |
| `resolverType` | `string` | no |
| `impactedServices` | `List<string>` | no |
| `caseAttachments` | `List<CaseAttachmentAttributes>` | no |
| `closedDate` | `timestamp` | no |
| `caseMetadata` | `List<CaseMetadataEntry>` | no |

## GetCaseAttachmentDownloadUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `attachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentPresignedUrl` | `string` | yes |

## GetCaseAttachmentUploadUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `fileName` | `string` | yes |
| `contentLength` | `long` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentPresignedUrl` | `string` | yes |

## GetMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |
| `accountId` | `string` | no |
| `region` | `string` | no |
| `membershipName` | `string` | no |
| `membershipArn` | `string` | no |
| `membershipStatus` | `string` | no |
| `membershipActivationTimestamp` | `timestamp` | no |
| `membershipDeactivationTimestamp` | `timestamp` | no |
| `customerType` | `string` | no |
| `numberOfAccountsCovered` | `long` | no |
| `incidentResponseTeam` | `List<IncidentResponder>` | no |
| `optInFeatures` | `List<OptInFeature>` | no |
| `membershipAccountsConfigurations` | `MembershipAccountsConfigurations` | no |

## ListCaseEdits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `caseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<CaseEditItem>` | no |
| `total` | `integer` | no |

## ListCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ListCasesItem>` | no |
| `total` | `long` | no |

## ListComments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `caseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ListCommentsItem>` | no |
| `total` | `integer` | no |

## ListInvestigations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `caseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `investigationActions` | `List<InvestigationAction>` | yes |

## ListMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ListMembershipItem>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## SendFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `resultId` | `string` | yes |
| `usefulness` | `string` | yes |
| `comment` | `string` | no |

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


## UpdateCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `title` | `string` | no |
| `description` | `string` | no |
| `reportedIncidentStartDate` | `timestamp` | no |
| `actualIncidentStartDate` | `timestamp` | no |
| `engagementType` | `string` | no |
| `watchersToAdd` | `List<Watcher>` | no |
| `watchersToDelete` | `List<Watcher>` | no |
| `threatActorIpAddressesToAdd` | `List<ThreatActorIp>` | no |
| `threatActorIpAddressesToDelete` | `List<ThreatActorIp>` | no |
| `impactedServicesToAdd` | `List<string>` | no |
| `impactedServicesToDelete` | `List<string>` | no |
| `impactedAwsRegionsToAdd` | `List<ImpactedAwsRegion>` | no |
| `impactedAwsRegionsToDelete` | `List<ImpactedAwsRegion>` | no |
| `impactedAccountsToAdd` | `List<string>` | no |
| `impactedAccountsToDelete` | `List<string>` | no |
| `caseMetadata` | `List<CaseMetadataEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCaseComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `commentId` | `string` | yes |
| `body` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |
| `body` | `string` | no |

## UpdateCaseStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `caseStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseStatus` | `string` | no |

## UpdateMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipId` | `string` | yes |
| `membershipName` | `string` | no |
| `incidentResponseTeam` | `List<IncidentResponder>` | no |
| `optInFeatures` | `List<OptInFeature>` | no |
| `membershipAccountsConfigurationsUpdate` | `MembershipAccountsConfigurationsUpdate` | no |
| `undoMembershipCancellation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResolverType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `resolverType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `caseStatus` | `string` | no |
| `resolverType` | `string` | no |

