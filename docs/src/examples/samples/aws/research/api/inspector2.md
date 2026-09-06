# Inspector2

API version: 2020-06-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/inspector2/2020-06-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

## BatchAssociateCodeSecurityScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associateConfigurationRequests` | `List<AssociateConfigurationRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failedAssociations` | `List<FailedAssociationResult>` | no |
| `successfulAssociations` | `List<SuccessfulAssociationResult>` | no |

## BatchDisassociateCodeSecurityScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `disassociateConfigurationRequests` | `List<DisassociateConfigurationRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failedAssociations` | `List<FailedAssociationResult>` | no |
| `successfulAssociations` | `List<SuccessfulAssociationResult>` | no |

## BatchGetAccountStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<AccountState>` | yes |
| `failedAccounts` | `List<FailedAccount>` | no |

## BatchGetCodeSnippet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeSnippetResults` | `List<CodeSnippetResult>` | no |
| `errors` | `List<CodeSnippetError>` | no |

## BatchGetFindingDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingDetails` | `List<FindingDetail>` | no |
| `errors` | `List<FindingDetailsError>` | no |

## BatchGetFreeTrialInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<FreeTrialAccountInfo>` | yes |
| `failedAccounts` | `List<FreeTrialInfoError>` | yes |

## BatchGetMemberEc2DeepInspectionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<MemberAccountEc2DeepInspectionStatusState>` | no |
| `failedAccountIds` | `List<FailedMemberAccountEc2DeepInspectionStatusState>` | no |

## BatchUpdateMemberEc2DeepInspectionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<MemberAccountEc2DeepInspectionStatus>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<MemberAccountEc2DeepInspectionStatusState>` | no |
| `failedAccountIds` | `List<FailedMemberAccountEc2DeepInspectionStatusState>` | no |

## CancelFindingsReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |

## CancelSbomExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

## CreateCisScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanName` | `string` | yes |
| `securityLevel` | `string` | yes |
| `schedule` | `Schedule` | yes |
| `targets` | `CreateCisTargets` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | no |

## CreateCodeSecurityIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `details` | `CreateIntegrationDetail` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `status` | `string` | yes |
| `authorizationUrl` | `string` | no |

## CreateCodeSecurityScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `level` | `string` | yes |
| `configuration` | `CodeSecurityScanConfiguration` | yes |
| `scopeSettings` | `ScopeSettings` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `provider` | `string` | yes |
| `description` | `string` | no |
| `providerDetail` | `ProviderDetailCreate` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |

## CreateFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `description` | `string` | no |
| `filterCriteria` | `FilterCriteria` | yes |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## CreateFindingsReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterCriteria` | `FilterCriteria` | no |
| `reportFormat` | `string` | yes |
| `s3Destination` | `Destination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

## CreateSbomExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceFilterCriteria` | `ResourceFilterCriteria` | no |
| `reportFormat` | `string` | yes |
| `s3Destination` | `Destination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

## DeleteCisScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |

## DeleteCodeSecurityIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | no |

## DeleteCodeSecurityScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | no |

## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## DescribeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnable` | `AutoEnable` | no |
| `maxAccountLimitReached` | `boolean` | no |

## Disable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `resourceTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<Account>` | yes |
| `failedAccounts` | `List<FailedAccount>` | no |

## DisableDelegatedAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegatedAdminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegatedAdminAccountId` | `string` | yes |

## DisassociateMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

## Enable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `resourceTypes` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<Account>` | yes |
| `failedAccounts` | `List<FailedAccount>` | no |

## EnableDelegatedAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegatedAdminAccountId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegatedAdminAccountId` | `string` | yes |

## GetCisScanReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanArn` | `string` | yes |
| `targetAccounts` | `List<string>` | no |
| `reportFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `url` | `string` | no |
| `status` | `string` | no |

## GetCisScanResultDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanArn` | `string` | yes |
| `targetResourceId` | `string` | yes |
| `accountId` | `string` | yes |
| `filterCriteria` | `CisScanResultDetailsFilterCriteria` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanResultDetails` | `List<CisScanResultDetails>` | no |
| `nextToken` | `string` | no |

## GetClustersForImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `ClusterForImageFilterCriteria` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `List<ClusterInformation>` | yes |
| `nextToken` | `string` | no |

## GetCodeSecurityIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | yes |
| `createdOn` | `timestamp` | yes |
| `lastUpdateOn` | `timestamp` | yes |
| `authorizationUrl` | `string` | no |
| `tags` | `Map<string>` | no |

## GetCodeSecurityScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resource` | `CodeSecurityResource` | yes |
| `scanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanId` | `string` | no |
| `resource` | `CodeSecurityResource` | no |
| `accountId` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `lastCommitId` | `string` | no |

## GetCodeSecurityScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | no |
| `name` | `string` | no |
| `configuration` | `CodeSecurityScanConfiguration` | no |
| `level` | `string` | no |
| `scopeSettings` | `ScopeSettings` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrConfiguration` | `EcrConfigurationState` | no |
| `ec2Configuration` | `Ec2ConfigurationState` | no |

## GetDelegatedAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegatedAdmin` | `DelegatedAdmin` | no |

## GetEc2DeepInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packagePaths` | `List<string>` | no |
| `orgPackagePaths` | `List<string>` | no |
| `status` | `string` | no |
| `errorMessage` | `string` | no |

## GetEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanType` | `string` | yes |
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKeyId` | `string` | yes |

## GetFindingsReportStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |
| `status` | `string` | no |
| `errorCode` | `string` | no |
| `errorMessage` | `string` | no |
| `destination` | `Destination` | no |
| `filterCriteria` | `FilterCriteria` | no |

## GetMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `member` | `Member` | no |

## GetSbomExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |
| `format` | `string` | no |
| `status` | `string` | no |
| `errorCode` | `string` | no |
| `errorMessage` | `string` | no |
| `s3Destination` | `Destination` | no |
| `filterCriteria` | `ResourceFilterCriteria` | no |

## ListAccountPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissions` | `List<Permission>` | yes |
| `nextToken` | `string` | no |

## ListCisScanConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterCriteria` | `ListCisScanConfigurationsFilterCriteria` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurations` | `List<CisScanConfiguration>` | no |
| `nextToken` | `string` | no |

## ListCisScanResultsAggregatedByChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanArn` | `string` | yes |
| `filterCriteria` | `CisScanResultsAggregatedByChecksFilterCriteria` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkAggregations` | `List<CisCheckAggregation>` | no |
| `nextToken` | `string` | no |

## ListCisScanResultsAggregatedByTargetResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanArn` | `string` | yes |
| `filterCriteria` | `CisScanResultsAggregatedByTargetResourceFilterCriteria` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetResourceAggregations` | `List<CisTargetResourceAggregation>` | no |
| `nextToken` | `string` | no |

## ListCisScans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterCriteria` | `ListCisScansFilterCriteria` | no |
| `detailLevel` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scans` | `List<CisScan>` | no |
| `nextToken` | `string` | no |

## ListCodeSecurityIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrations` | `List<CodeSecurityIntegrationSummary>` | no |
| `nextToken` | `string` | no |

## ListCodeSecurityScanConfigurationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associations` | `List<CodeSecurityScanConfigurationAssociationSummary>` | no |
| `nextToken` | `string` | no |

## ListCodeSecurityScanConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<CodeSecurityScanConfigurationSummary>` | no |
| `nextToken` | `string` | no |

## ListConnectorScanConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `awsConfigConnectorArns` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurations` | `List<ConnectorScanConfigurationItem>` | yes |
| `nextToken` | `string` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filterCriteria` | `ConnectorFilterCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Connector>` | yes |
| `nextToken` | `string` | no |

## ListCoverage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filterCriteria` | `CoverageFilterCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `coveredResources` | `List<CoveredResource>` | no |

## ListCoverageStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterCriteria` | `CoverageFilterCriteria` | no |
| `groupBy` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `countsByGroup` | `List<Counts>` | no |
| `totalCounts` | `long` | yes |
| `nextToken` | `string` | no |

## ListDelegatedAdminAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delegatedAdminAccounts` | `List<DelegatedAdminAccount>` | no |
| `nextToken` | `string` | no |

## ListFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arns` | `List<string>` | no |
| `action` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | yes |
| `nextToken` | `string` | no |

## ListFindingAggregations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aggregationType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `accountIds` | `List<StringFilter>` | no |
| `aggregationRequest` | `AggregationRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aggregationType` | `string` | yes |
| `responses` | `List<AggregationResponse>` | no |
| `nextToken` | `string` | no |

## ListFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filterCriteria` | `FilterCriteria` | no |
| `sortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `findings` | `List<Finding>` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlyAssociated` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<Member>` | no |
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

## ListUsageTotals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `totals` | `List<UsageTotal>` | no |

## ResetEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanType` | `string` | yes |
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchVulnerabilities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterCriteria` | `SearchVulnerabilitiesFilterCriteria` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vulnerabilities` | `List<Vulnerability>` | yes |
| `nextToken` | `string` | no |

## SendCisSessionHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanJobId` | `string` | yes |
| `sessionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendCisSessionTelemetry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanJobId` | `string` | yes |
| `sessionToken` | `string` | yes |
| `messages` | `List<CisSessionMessage>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCisSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanJobId` | `string` | yes |
| `message` | `StartCisSessionMessage` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCodeSecurityScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `resource` | `CodeSecurityResource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanId` | `string` | no |
| `status` | `string` | no |

## StopCisSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanJobId` | `string` | yes |
| `sessionToken` | `string` | yes |
| `message` | `StopCisSessionMessage` | yes |

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


## UpdateCisScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |
| `scanName` | `string` | no |
| `securityLevel` | `string` | no |
| `schedule` | `Schedule` | no |
| `targets` | `UpdateCisTargets` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |

## UpdateCodeSecurityIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `details` | `UpdateIntegrationDetails` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `status` | `string` | yes |

## UpdateCodeSecurityScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | yes |
| `configuration` | `CodeSecurityScanConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanConfigurationArn` | `string` | no |

## UpdateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |
| `ecrConfiguration` | `EcrConfiguration` | no |
| `ec2Configuration` | `Ec2Configuration` | no |
| `updateConfigurationInheritance` | `UpdateConfigurationInheritance` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |
| `description` | `string` | no |
| `providerDetail` | `ProviderDetailUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |

## UpdateConnectorScanConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `awsConfigConnectorArn` | `string` | yes |
| `scanConfiguration` | `ConnectorScanConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEc2DeepInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activateDeepInspection` | `boolean` | no |
| `packagePaths` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packagePaths` | `List<string>` | no |
| `orgPackagePaths` | `List<string>` | no |
| `status` | `string` | no |
| `errorMessage` | `string` | no |

## UpdateEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKeyId` | `string` | yes |
| `scanType` | `string` | yes |
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | no |
| `description` | `string` | no |
| `filterCriteria` | `FilterCriteria` | no |
| `name` | `string` | no |
| `filterArn` | `string` | yes |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## UpdateOrgEc2DeepInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `orgPackagePaths` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnable` | `AutoEnable` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnable` | `AutoEnable` | yes |

