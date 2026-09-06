# Firewall Management Service

API version: 2018-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/fms/2018-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccount` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateThirdPartyFirewall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewall` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewallStatus` | `string` | no |

## BatchAssociateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetIdentifier` | `string` | yes |
| `Items` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetIdentifier` | `string` | yes |
| `FailedItems` | `List<FailedItem>` | yes |

## BatchDisassociateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetIdentifier` | `string` | yes |
| `Items` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetIdentifier` | `string` | yes |
| `FailedItems` | `List<FailedItem>` | yes |

## DeleteAppsList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotificationChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `DeleteAllPolicyResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProtocolsList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateThirdPartyFirewall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewall` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewallStatus` | `string` | no |

## GetAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccount` | `string` | no |
| `RoleStatus` | `string` | no |

## GetAdminScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccount` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminScope` | `AdminScope` | no |
| `Status` | `string` | no |

## GetAppsList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListId` | `string` | yes |
| `DefaultList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppsList` | `AppsListData` | no |
| `AppsListArn` | `string` | no |

## GetComplianceDetail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `MemberAccount` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyComplianceDetail` | `PolicyComplianceDetail` | no |

## GetNotificationChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnsTopicArn` | `string` | no |
| `SnsRoleName` | `string` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |
| `PolicyArn` | `string` | no |

## GetProtectionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `MemberAccountId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccountId` | `string` | no |
| `ServiceType` | `string` | no |
| `Data` | `string` | no |
| `NextToken` | `string` | no |

## GetProtocolsList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListId` | `string` | yes |
| `DefaultList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtocolsList` | `ProtocolsListData` | no |
| `ProtocolsListArn` | `string` | no |

## GetResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSet` | `ResourceSet` | yes |
| `ResourceSetArn` | `string` | yes |

## GetThirdPartyFirewallAssociationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewall` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewallStatus` | `string` | no |
| `MarketplaceOnboardingStatus` | `string` | no |

## GetViolationDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `MemberAccount` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViolationDetail` | `ViolationDetail` | no |

## ListAdminAccountsForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccounts` | `List<AdminAccountSummary>` | no |
| `NextToken` | `string` | no |

## ListAdminsManagingAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccounts` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListAppsLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultLists` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppsLists` | `List<AppsListDataSummary>` | no |
| `NextToken` | `string` | no |

## ListComplianceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyComplianceStatusList` | `List<PolicyComplianceStatus>` | no |
| `NextToken` | `string` | no |

## ListDiscoveredResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberAccountIds` | `List<string>` | yes |
| `ResourceType` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DiscoveredResource>` | no |
| `NextToken` | `string` | no |

## ListMemberAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberAccounts` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyList` | `List<PolicySummary>` | no |
| `NextToken` | `string` | no |

## ListProtocolsLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultLists` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtocolsLists` | `List<ProtocolsListDataSummary>` | no |
| `NextToken` | `string` | no |

## ListResourceSetResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Resource>` | yes |
| `NextToken` | `string` | no |

## ListResourceSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSets` | `List<ResourceSetSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ListThirdPartyFirewallFirewallPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewall` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThirdPartyFirewallFirewallPolicies` | `List<ThirdPartyFirewallFirewallPolicy>` | no |
| `NextToken` | `string` | no |

## PutAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccount` | `string` | yes |
| `AdminScope` | `AdminScope` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAppsList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppsList` | `AppsListData` | yes |
| `TagList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppsList` | `AppsListData` | no |
| `AppsListArn` | `string` | no |

## PutNotificationChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnsTopicArn` | `string` | yes |
| `SnsRoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | yes |
| `TagList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |
| `PolicyArn` | `string` | no |

## PutProtocolsList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtocolsList` | `ProtocolsListData` | yes |
| `TagList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtocolsList` | `ProtocolsListData` | no |
| `ProtocolsListArn` | `string` | no |

## PutResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSet` | `ResourceSet` | yes |
| `TagList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSet` | `ResourceSet` | yes |
| `ResourceSetArn` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagList` | `List<Tag>` | yes |

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


