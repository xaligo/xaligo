# AWS Organizations

API version: 2016-11-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/organizations/2016-11-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HandshakeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## AttachPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `TargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HandshakeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## CloseAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Email` | `string` | yes |
| `AccountName` | `string` | yes |
| `RoleName` | `string` | no |
| `IamUserAccessToBilling` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateAccountStatus` | `CreateAccountStatus` | no |

## CreateGovCloudAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Email` | `string` | yes |
| `AccountName` | `string` | yes |
| `RoleName` | `string` | no |
| `IamUserAccessToBilling` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateAccountStatus` | `CreateAccountStatus` | no |

## CreateOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureSet` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Organization` | `Organization` | no |

## CreateOrganizationalUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParentId` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnit` | `OrganizationalUnit` | no |

## CreatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | yes |
| `Description` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## DeclineHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HandshakeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## DeleteOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOrganizationalUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterDelegatedAdministrator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ServicePrincipal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Account` | `Account` | no |

## DescribeCreateAccountStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateAccountRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateAccountStatus` | `CreateAccountStatus` | no |

## DescribeEffectivePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyType` | `string` | yes |
| `TargetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EffectivePolicy` | `EffectivePolicy` | no |

## DescribeHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HandshakeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## DescribeOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Organization` | `Organization` | no |

## DescribeOrganizationalUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnit` | `OrganizationalUnit` | no |

## DescribePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## DescribeResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `ResourcePolicy` | no |

## DescribeResponsibilityTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponsibilityTransfer` | `ResponsibilityTransfer` | no |

## DetachPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `TargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableAWSServiceAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisablePolicyType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RootId` | `string` | yes |
| `PolicyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Root` | `Root` | no |

## EnableAWSServiceAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableAllFeatures

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## EnablePolicyType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RootId` | `string` | yes |
| `PolicyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Root` | `Root` | no |

## InviteAccountToOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Target` | `HandshakeParty` | yes |
| `Notes` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## InviteOrganizationToTransferResponsibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `Target` | `HandshakeParty` | yes |
| `Notes` | `string` | no |
| `StartTimestamp` | `timestamp` | yes |
| `SourceName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshake` | `Handshake` | no |

## LeaveOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListAWSServiceAccessForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnabledServicePrincipals` | `List<EnabledServicePrincipal>` | no |
| `NextToken` | `string` | no |

## ListAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accounts` | `List<Account>` | no |
| `NextToken` | `string` | no |

## ListAccountsForParent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParentId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accounts` | `List<Account>` | no |
| `NextToken` | `string` | no |

## ListAccountsWithInvalidEffectivePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyType` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accounts` | `List<Account>` | no |
| `PolicyType` | `string` | no |
| `NextToken` | `string` | no |

## ListChildren

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParentId` | `string` | yes |
| `ChildType` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Children` | `List<Child>` | no |
| `NextToken` | `string` | no |

## ListCreateAccountStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `States` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateAccountStatuses` | `List<CreateAccountStatus>` | no |
| `NextToken` | `string` | no |

## ListDelegatedAdministrators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipal` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegatedAdministrators` | `List<DelegatedAdministrator>` | no |
| `NextToken` | `string` | no |

## ListDelegatedServicesForAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegatedServices` | `List<DelegatedService>` | no |
| `NextToken` | `string` | no |

## ListEffectivePolicyValidationErrors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `PolicyType` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `PolicyType` | `string` | no |
| `Path` | `string` | no |
| `EvaluationTimestamp` | `timestamp` | no |
| `NextToken` | `string` | no |
| `EffectivePolicyValidationErrors` | `List<EffectivePolicyValidationError>` | no |

## ListHandshakesForAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `HandshakeFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshakes` | `List<Handshake>` | no |
| `NextToken` | `string` | no |

## ListHandshakesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `HandshakeFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Handshakes` | `List<Handshake>` | no |
| `NextToken` | `string` | no |

## ListInboundResponsibilityTransfers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `Id` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponsibilityTransfers` | `List<ResponsibilityTransfer>` | no |
| `NextToken` | `string` | no |

## ListOrganizationalUnitsForParent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParentId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnits` | `List<OrganizationalUnit>` | no |
| `NextToken` | `string` | no |

## ListOutboundResponsibilityTransfers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponsibilityTransfers` | `List<ResponsibilityTransfer>` | no |
| `NextToken` | `string` | no |

## ListParents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChildId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parents` | `List<Parent>` | no |
| `NextToken` | `string` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `List<PolicySummary>` | no |
| `NextToken` | `string` | no |

## ListPoliciesForTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetId` | `string` | yes |
| `Filter` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `List<PolicySummary>` | no |
| `NextToken` | `string` | no |

## ListRoots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Roots` | `List<Root>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## ListTargetsForPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Targets` | `List<PolicyTargetSummary>` | no |
| `NextToken` | `string` | no |

## MoveAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `SourceParentId` | `string` | yes |
| `DestinationParentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `ResourcePolicy` | no |

## RegisterDelegatedAdministrator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ServicePrincipal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveAccountFromOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateResponsibilityTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `EndTimestamp` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponsibilityTransfer` | `ResponsibilityTransfer` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOrganizationalUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnitId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationalUnit` | `OrganizationalUnit` | no |

## UpdatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Content` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## UpdateResponsibilityTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponsibilityTransfer` | `ResponsibilityTransfer` | no |

