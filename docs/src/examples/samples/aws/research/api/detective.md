# Amazon Detective

API version: 2018-10-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/detective/2018-10-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchGetGraphMemberDatasources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberDatasources` | `List<MembershipDatasources>` | no |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | no |

## BatchGetMembershipDatasources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MembershipDatasources` | `List<MembershipDatasources>` | no |
| `UnprocessedGraphs` | `List<UnprocessedGraph>` | no |

## CreateGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | no |

## CreateMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `Message` | `string` | no |
| `DisableEmailNotification` | `boolean` | no |
| `Accounts` | `List<Account>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Members` | `List<MemberDetail>` | no |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | no |

## DeleteGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | no |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | no |

## DescribeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnable` | `boolean` | no |

## DisableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetInvestigation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `InvestigationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | no |
| `InvestigationId` | `string` | no |
| `EntityArn` | `string` | no |
| `EntityType` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `ScopeStartTime` | `timestamp` | no |
| `ScopeEndTime` | `timestamp` | no |
| `Status` | `string` | no |
| `Severity` | `string` | no |
| `State` | `string` | no |

## GetMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberDetails` | `List<MemberDetail>` | no |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | no |

## ListDatasourcePackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasourcePackages` | `Map<DatasourcePackageIngestDetail>` | no |
| `NextToken` | `string` | no |

## ListGraphs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphList` | `List<Graph>` | no |
| `NextToken` | `string` | no |

## ListIndicators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `InvestigationId` | `string` | yes |
| `IndicatorType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | no |
| `InvestigationId` | `string` | no |
| `NextToken` | `string` | no |
| `Indicators` | `List<Indicator>` | no |

## ListInvestigations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `SortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvestigationDetails` | `List<InvestigationDetail>` | no |
| `NextToken` | `string` | no |

## ListInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Invitations` | `List<MemberDetail>` | no |
| `NextToken` | `string` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberDetails` | `List<MemberDetail>` | no |
| `NextToken` | `string` | no |

## ListOrganizationAdminAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Administrators` | `List<Administrator>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## RejectInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartInvestigation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `EntityArn` | `string` | yes |
| `ScopeStartTime` | `timestamp` | yes |
| `ScopeEndTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvestigationId` | `string` | no |

## StartMonitoringMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateDatasourcePackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `DatasourcePackages` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInvestigationState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `InvestigationId` | `string` | yes |
| `State` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GraphArn` | `string` | yes |
| `AutoEnable` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


