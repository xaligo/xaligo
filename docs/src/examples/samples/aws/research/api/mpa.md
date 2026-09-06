# AWS Multi-party Approval

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mpa/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateApprovalTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ApprovalStrategy` | `ApprovalStrategy` | yes |
| `Approvers` | `List<ApprovalTeamRequestApprover>` | yes |
| `Description` | `string` | yes |
| `Policies` | `List<PolicyReference>` | yes |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |

## CreateIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentitySourceParameters` | `IdentitySourceParameters` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentitySourceType` | `string` | no |
| `IdentitySourceArn` | `string` | no |
| `CreationTime` | `timestamp` | no |

## DeleteIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentitySourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInactiveApprovalTeamVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApprovalTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | no |
| `ApprovalStrategy` | `ApprovalStrategyResponse` | no |
| `NumberOfApprovers` | `integer` | no |
| `Approvers` | `List<GetApprovalTeamResponseApprover>` | no |
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `StatusCode` | `string` | no |
| `StatusMessage` | `string` | no |
| `UpdateSessionArn` | `string` | no |
| `VersionId` | `string` | no |
| `Policies` | `List<PolicyReference>` | no |
| `LastUpdateTime` | `timestamp` | no |
| `PendingUpdate` | `PendingUpdate` | no |

## GetIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentitySourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentitySourceType` | `string` | no |
| `IdentitySourceParameters` | `IdentitySourceParametersForGet` | no |
| `IdentitySourceArn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `Status` | `string` | no |
| `StatusCode` | `string` | no |
| `StatusMessage` | `string` | no |

## GetPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyVersion` | `PolicyVersion` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `PolicyType` | `string` | yes |
| `PolicyVersionArn` | `string` | no |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionArn` | `string` | no |
| `ApprovalTeamArn` | `string` | no |
| `ApprovalTeamName` | `string` | no |
| `ProtectedResourceArn` | `string` | no |
| `ApprovalStrategy` | `ApprovalStrategyResponse` | no |
| `NumberOfApprovers` | `integer` | no |
| `InitiationTime` | `timestamp` | no |
| `ExpirationTime` | `timestamp` | no |
| `CompletionTime` | `timestamp` | no |
| `Description` | `string` | no |
| `Metadata` | `Map<string>` | no |
| `Status` | `string` | no |
| `StatusCode` | `string` | no |
| `StatusMessage` | `string` | no |
| `ExecutionStatus` | `string` | no |
| `ActionName` | `string` | no |
| `RequesterServicePrincipal` | `string` | no |
| `RequesterPrincipalArn` | `string` | no |
| `RequesterAccountId` | `string` | no |
| `RequesterRegion` | `string` | no |
| `RequesterComment` | `string` | no |
| `ActionCompletionStrategy` | `string` | no |
| `ApproverResponses` | `List<GetSessionResponseApproverResponse>` | no |
| `AdditionalSecurityRequirements` | `List<string>` | no |

## ListApprovalTeams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ApprovalTeams` | `List<ListApprovalTeamsResponseApprovalTeam>` | no |

## ListIdentitySources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IdentitySources` | `List<IdentitySourceForList>` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Policies` | `List<Policy>` | no |

## ListPolicyVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PolicyVersions` | `List<PolicyVersionSummary>` | no |

## ListResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ResourcePolicies` | `List<ListResourcePoliciesResponseResourcePolicy>` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApprovalTeamArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Sessions` | `List<ListSessionsResponseSession>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartActiveApprovalTeamDeletion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PendingWindowDays` | `integer` | no |
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionCompletionTime` | `timestamp` | no |
| `DeletionStartTime` | `timestamp` | no |

## StartApprovalTeamBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ApproverIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineSessionArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateApprovalTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApprovalStrategy` | `ApprovalStrategy` | no |
| `Approvers` | `List<ApprovalTeamRequestApprover>` | no |
| `Description` | `string` | no |
| `Arn` | `string` | yes |
| `UpdateActions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VersionId` | `string` | no |

