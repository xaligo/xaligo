# AWS Single Sign-On Admin

API version: 2020-07-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sso-admin/2020-07-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## AttachCustomerManagedPolicyReferenceToPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `CustomerManagedPolicyReference` | `CustomerManagedPolicyReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachManagedPolicyToPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `ManagedPolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAccountAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `TargetId` | `string` | yes |
| `TargetType` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `PrincipalType` | `string` | yes |
| `PrincipalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignmentCreationStatus` | `AccountAssignmentOperationStatus` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `ApplicationProviderArn` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `PortalOptions` | `PortalOptions` | no |
| `Tags` | `List<Tag>` | no |
| `Status` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | no |
| `InstanceArn` | `string` | no |
| `IdentityStoreArn` | `string` | no |

## CreateApplicationAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `PrincipalId` | `string` | yes |
| `PrincipalType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | no |

## CreateInstanceAccessControlAttributeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `InstanceAccessControlAttributeConfiguration` | `InstanceAccessControlAttributeConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreatePermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `InstanceArn` | `string` | yes |
| `SessionDuration` | `string` | no |
| `RelayState` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionSet` | `PermissionSet` | no |

## CreateTrustedTokenIssuer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `Name` | `string` | yes |
| `TrustedTokenIssuerType` | `string` | yes |
| `TrustedTokenIssuerConfiguration` | `TrustedTokenIssuerConfiguration` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedTokenIssuerArn` | `string` | no |

## DeleteAccountAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `TargetId` | `string` | yes |
| `TargetType` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `PrincipalType` | `string` | yes |
| `PrincipalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignmentDeletionStatus` | `AccountAssignmentOperationStatus` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationAccessScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `Scope` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `PrincipalId` | `string` | yes |
| `PrincipalType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationAuthenticationMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `AuthenticationMethodType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `GrantType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInlinePolicyFromPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstanceAccessControlAttributeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePermissionsBoundaryFromPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrustedTokenIssuer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedTokenIssuerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountAssignmentCreationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `AccountAssignmentCreationRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignmentCreationStatus` | `AccountAssignmentOperationStatus` | no |

## DescribeAccountAssignmentDeletionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `AccountAssignmentDeletionRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignmentDeletionStatus` | `AccountAssignmentOperationStatus` | no |

## DescribeApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | no |
| `ApplicationProviderArn` | `string` | no |
| `Name` | `string` | no |
| `ApplicationAccount` | `string` | no |
| `InstanceArn` | `string` | no |
| `IdentityStoreArn` | `string` | no |
| `Status` | `string` | no |
| `PortalOptions` | `PortalOptions` | no |
| `Description` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `CreatedFrom` | `string` | no |

## DescribeApplicationAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `PrincipalId` | `string` | yes |
| `PrincipalType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrincipalType` | `string` | no |
| `PrincipalId` | `string` | no |
| `ApplicationArn` | `string` | no |

## DescribeApplicationProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationProviderArn` | `string` | yes |
| `FederationProtocol` | `string` | no |
| `DisplayData` | `DisplayData` | no |
| `ResourceServerConfig` | `ResourceServerConfig` | no |

## DescribeInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | no |
| `IdentityStoreId` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `Name` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `EncryptionConfigurationDetails` | `EncryptionConfigurationDetails` | no |
| `PermissionSetsEnabled` | `boolean` | no |

## DescribeInstanceAccessControlAttributeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `InstanceAccessControlAttributeConfiguration` | `InstanceAccessControlAttributeConfiguration` | no |

## DescribePermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionSet` | `PermissionSet` | no |

## DescribePermissionSetProvisioningStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `ProvisionPermissionSetRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionSetProvisioningStatus` | `PermissionSetProvisioningStatus` | no |

## DescribeRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionName` | `string` | no |
| `Status` | `string` | no |
| `AddedDate` | `timestamp` | no |
| `IsPrimaryRegion` | `boolean` | no |

## DescribeTrustedTokenIssuer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedTokenIssuerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedTokenIssuerArn` | `string` | no |
| `Name` | `string` | no |
| `TrustedTokenIssuerType` | `string` | no |
| `TrustedTokenIssuerConfiguration` | `TrustedTokenIssuerConfiguration` | no |

## DetachCustomerManagedPolicyReferenceFromPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `CustomerManagedPolicyReference` | `CustomerManagedPolicyReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachManagedPolicyFromPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `ManagedPolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplicationAccessScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `Scope` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `AuthorizedTargets` | `List<string>` | no |

## GetApplicationAssignmentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentRequired` | `boolean` | yes |

## GetApplicationAuthenticationMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `AuthenticationMethodType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationMethod` | `AuthenticationMethod` | no |

## GetApplicationGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `GrantType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grant` | `Grant` | yes |

## GetApplicationSessionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserBackgroundSessionApplicationStatus` | `string` | no |

## GetInlinePolicyForPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InlinePolicy` | `string` | no |

## GetPermissionsBoundaryForPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionsBoundary` | `PermissionsBoundary` | no |

## ListAccountAssignmentCreationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filter` | `OperationStatusFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignmentsCreationStatus` | `List<AccountAssignmentOperationStatusMetadata>` | no |
| `NextToken` | `string` | no |

## ListAccountAssignmentDeletionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filter` | `OperationStatusFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignmentsDeletionStatus` | `List<AccountAssignmentOperationStatusMetadata>` | no |
| `NextToken` | `string` | no |

## ListAccountAssignments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `AccountId` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignments` | `List<AccountAssignment>` | no |
| `NextToken` | `string` | no |

## ListAccountAssignmentsForPrincipal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PrincipalId` | `string` | yes |
| `PrincipalType` | `string` | yes |
| `Filter` | `ListAccountAssignmentsFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssignments` | `List<AccountAssignmentForPrincipal>` | no |
| `NextToken` | `string` | no |

## ListAccountsForProvisionedPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `ProvisioningStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListApplicationAccessScopes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scopes` | `List<ScopeDetails>` | yes |
| `NextToken` | `string` | no |

## ListApplicationAssignments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationAssignments` | `List<ApplicationAssignment>` | no |
| `NextToken` | `string` | no |

## ListApplicationAssignmentsForPrincipal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PrincipalId` | `string` | yes |
| `PrincipalType` | `string` | yes |
| `Filter` | `ListApplicationAssignmentsFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationAssignments` | `List<ApplicationAssignmentForPrincipal>` | no |
| `NextToken` | `string` | no |

## ListApplicationAuthenticationMethods

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationMethods` | `List<AuthenticationMethodItem>` | no |
| `NextToken` | `string` | no |

## ListApplicationGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grants` | `List<GrantItem>` | yes |
| `NextToken` | `string` | no |

## ListApplicationProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationProviders` | `List<ApplicationProvider>` | no |
| `NextToken` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filter` | `ListApplicationsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<Application>` | no |
| `NextToken` | `string` | no |

## ListCustomerManagedPolicyReferencesInPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerManagedPolicyReferences` | `List<CustomerManagedPolicyReference>` | no |
| `NextToken` | `string` | no |

## ListInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<InstanceMetadata>` | no |
| `NextToken` | `string` | no |

## ListManagedPoliciesInPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedManagedPolicies` | `List<AttachedManagedPolicy>` | no |
| `NextToken` | `string` | no |

## ListPermissionSetProvisioningStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filter` | `OperationStatusFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionSetsProvisioningStatus` | `List<PermissionSetProvisioningStatusMetadata>` | no |
| `NextToken` | `string` | no |

## ListPermissionSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionSets` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListPermissionSetsProvisionedToAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `AccountId` | `string` | yes |
| `ProvisioningStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PermissionSets` | `List<string>` | no |

## ListRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<RegionMetadata>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | no |
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## ListTrustedTokenIssuers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedTokenIssuers` | `List<TrustedTokenIssuerMetadata>` | no |
| `NextToken` | `string` | no |

## ProvisionPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `TargetId` | `string` | no |
| `TargetType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PermissionSetProvisioningStatus` | `PermissionSetProvisioningStatus` | no |

## PutApplicationAccessScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `AuthorizedTargets` | `List<string>` | no |
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutApplicationAssignmentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `AssignmentRequired` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutApplicationAuthenticationMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `AuthenticationMethodType` | `string` | yes |
| `AuthenticationMethod` | `AuthenticationMethod` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutApplicationGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `GrantType` | `string` | yes |
| `Grant` | `Grant` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutApplicationSessionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `UserBackgroundSessionApplicationStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutInlinePolicyToPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `InlinePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutPermissionsBoundaryToPermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `PermissionsBoundary` | `PermissionsBoundary` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | no |
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | no |
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `PortalOptions` | `UpdateApplicationPortalOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `InstanceArn` | `string` | yes |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `PermissionSetsEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInstanceAccessControlAttributeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `InstanceAccessControlAttributeConfiguration` | `InstanceAccessControlAttributeConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePermissionSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `PermissionSetArn` | `string` | yes |
| `Description` | `string` | no |
| `SessionDuration` | `string` | no |
| `RelayState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTrustedTokenIssuer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedTokenIssuerArn` | `string` | yes |
| `Name` | `string` | no |
| `TrustedTokenIssuerConfiguration` | `TrustedTokenIssuerUpdateConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


