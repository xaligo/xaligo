# AWS Identity and Access Management

API version: 2010-05-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iam/2010-05-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptDelegationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AcquireRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `TemplateMinorVersion` | `integer` | no |
| `ReplacementValues` | `Map<ReplacementValueEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `Role` | yes |

## AddClientIDToOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |
| `ClientID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddRoleToInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |
| `RoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddUserToGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateDelegationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachRolePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachUserPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ChangePassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OldPassword` | `string` | yes |
| `NewPassword` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAccessKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKey` | `AccessKey` | yes |

## CreateAccountAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDelegationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerAccountId` | `string` | no |
| `Description` | `string` | yes |
| `Permissions` | `DelegationPermission` | yes |
| `RequestMessage` | `string` | no |
| `RequestorWorkflowId` | `string` | yes |
| `RedirectUrl` | `string` | no |
| `NotificationChannel` | `string` | yes |
| `SessionDuration` | `integer` | yes |
| `OnlySendByOwner` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsoleDeepLink` | `string` | no |
| `DelegationRequestId` | `string` | no |

## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | no |
| `GroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | yes |

## CreateInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |
| `Path` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfile` | `InstanceProfile` | yes |

## CreateLoginProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `Password` | `string` | no |
| `PasswordResetRequired` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoginProfile` | `LoginProfile` | yes |

## CreateOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | yes |
| `ClientIDList` | `List<string>` | no |
| `ThumbprintList` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## CreatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyName` | `string` | yes |
| `Path` | `string` | no |
| `PolicyDocument` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## CreatePolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `PolicyDocument` | `string` | yes |
| `SetAsDefault` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyVersion` | `PolicyVersion` | no |

## CreateRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | no |
| `RoleName` | `string` | yes |
| `AssumeRolePolicyDocument` | `string` | yes |
| `Description` | `string` | no |
| `MaxSessionDuration` | `integer` | no |
| `PermissionsBoundary` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `Role` | yes |

## CreateSAMLProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLMetadataDocument` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `AssertionEncryptionMode` | `string` | no |
| `AddPrivateKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## CreateServiceLinkedRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AWSServiceName` | `string` | yes |
| `Description` | `string` | no |
| `CustomSuffix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `Role` | no |

## CreateServiceSpecificCredential

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `ServiceName` | `string` | yes |
| `CredentialAgeDays` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSpecificCredential` | `ServiceSpecificCredential` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | no |
| `UserName` | `string` | yes |
| `PermissionsBoundary` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## CreateVirtualMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | no |
| `VirtualMFADeviceName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VirtualMFADevice` | `VirtualMFADevice` | yes |

## DeactivateMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `SerialNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `AccessKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccountAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccountPasswordPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoginProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRolePermissionsBoundary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRolePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSAMLProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSSHPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `SSHPublicKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServerCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceLinkedRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionTaskId` | `string` | yes |

## DeleteServiceSpecificCredential

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `ServiceSpecificCredentialId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSigningCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `CertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPermissionsBoundary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVirtualMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachRolePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachUserPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableOrganizationsRootCredentialsManagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `EnabledFeatures` | `List<string>` | no |

## DisableOrganizationsRootSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `EnabledFeatures` | `List<string>` | no |

## DisableOutboundWebIdentityFederation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `SerialNumber` | `string` | yes |
| `AuthenticationCode1` | `string` | yes |
| `AuthenticationCode2` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableOrganizationsRootCredentialsManagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `EnabledFeatures` | `List<string>` | no |

## EnableOrganizationsRootSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `EnabledFeatures` | `List<string>` | no |

## EnableOutboundWebIdentityFederation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IssuerIdentifier` | `string` | no |

## GenerateCredentialReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `Description` | `string` | no |

## GenerateOrganizationsAccessReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityPath` | `string` | yes |
| `OrganizationsPolicyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## GenerateServiceLastAccessedDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Granularity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## GetAccessKeyLastUsed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `AccessKeyLastUsed` | `AccessKeyLastUsed` | no |

## GetAccountAuthorizationDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `List<string>` | no |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserDetailList` | `List<UserDetail>` | no |
| `GroupDetailList` | `List<GroupDetail>` | no |
| `RoleDetailList` | `List<RoleDetail>` | no |
| `Policies` | `List<ManagedPolicyDetail>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## GetAccountPasswordPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PasswordPolicy` | `PasswordPolicy` | yes |

## GetAccountProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Properties` | `Map<string>` | no |

## GetAccountSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryMap` | `Map<integer>` | no |

## GetContextKeysForCustomPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyInputList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextKeyNames` | `List<string>` | no |

## GetContextKeysForPrincipalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicySourceArn` | `string` | yes |
| `PolicyInputList` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextKeyNames` | `List<string>` | no |

## GetCredentialReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `blob` | no |
| `ReportFormat` | `string` | no |
| `GeneratedTime` | `timestamp` | no |

## GetDelegationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequestId` | `string` | yes |
| `DelegationPermissionCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequest` | `DelegationRequest` | no |
| `PermissionCheckStatus` | `string` | no |
| `PermissionCheckResult` | `string` | no |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | yes |
| `Users` | `List<User>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## GetGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

## GetHumanReadableSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityArn` | `string` | yes |
| `Locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryContent` | `string` | no |
| `Locale` | `string` | no |
| `SummaryState` | `string` | no |

## GetInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfile` | `InstanceProfile` | yes |

## GetLoginProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoginProfile` | `LoginProfile` | yes |

## GetMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialNumber` | `string` | yes |
| `UserName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `SerialNumber` | `string` | yes |
| `EnableDate` | `timestamp` | no |
| `Certifications` | `Map<string>` | no |

## GetOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |
| `ClientIDList` | `List<string>` | no |
| `ThumbprintList` | `List<string>` | no |
| `CreateDate` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |

## GetOrganizationsAccessReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |
| `SortKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | yes |
| `JobCreationDate` | `timestamp` | yes |
| `JobCompletionDate` | `timestamp` | no |
| `NumberOfServicesAccessible` | `integer` | no |
| `NumberOfServicesNotAccessed` | `integer` | no |
| `AccessDetails` | `List<AccessDetail>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |
| `ErrorDetails` | `ErrorDetails` | no |

## GetOutboundWebIdentityFederationInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IssuerIdentifier` | `string` | no |
| `JwtVendingEnabled` | `boolean` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## GetPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyVersion` | `PolicyVersion` | no |

## GetRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `Role` | yes |

## GetRolePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

## GetRoleTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |
| `MinorVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleTemplateVersion` | `RoleTemplateVersion` | yes |

## GetSAMLProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderUUID` | `string` | no |
| `SAMLMetadataDocument` | `string` | no |
| `CreateDate` | `timestamp` | no |
| `ValidUntil` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |
| `AssertionEncryptionMode` | `string` | no |
| `PrivateKeyList` | `List<SAMLPrivateKey>` | no |

## GetSSHPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `SSHPublicKeyId` | `string` | yes |
| `Encoding` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SSHPublicKey` | `SSHPublicKey` | no |

## GetServerCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificate` | `ServerCertificate` | yes |

## GetServiceLastAccessedDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | yes |
| `JobType` | `string` | no |
| `JobCreationDate` | `timestamp` | yes |
| `ServicesLastAccessed` | `List<ServiceLastAccessed>` | yes |
| `JobCompletionDate` | `timestamp` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |
| `Error` | `ErrorDetails` | no |

## GetServiceLastAccessedDetailsWithEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `ServiceNamespace` | `string` | yes |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | yes |
| `JobCreationDate` | `timestamp` | yes |
| `JobCompletionDate` | `timestamp` | yes |
| `EntityDetailsList` | `List<EntityDetails>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |
| `Error` | `ErrorDetails` | no |

## GetServiceLinkedRoleDeletionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |
| `Reason` | `DeletionTaskFailureReasonType` | no |

## GetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | yes |

## GetUserPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

## ListAccessKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKeyMetadata` | `List<AccessKeyMetadata>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListAccountAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAliases` | `List<string>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListAttachedGroupPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedPolicies` | `List<AttachedPolicy>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListAttachedRolePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedPolicies` | `List<AttachedPolicy>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListAttachedUserPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedPolicies` | `List<AttachedPolicy>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListDelegationRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerId` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequests` | `List<DelegationRequest>` | no |
| `Marker` | `string` | no |
| `isTruncated` | `boolean` | no |

## ListEntitiesForPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `EntityFilter` | `string` | no |
| `PathPrefix` | `string` | no |
| `PolicyUsageFilter` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyGroups` | `List<PolicyGroup>` | no |
| `PolicyUsers` | `List<PolicyUser>` | no |
| `PolicyRoles` | `List<PolicyRole>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListGroupPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyNames` | `List<string>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<Group>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListGroupsForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<Group>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListInstanceProfileTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListInstanceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfiles` | `List<InstanceProfile>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListInstanceProfilesForRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfiles` | `List<InstanceProfile>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListMFADeviceTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialNumber` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListMFADevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MFADevices` | `List<MFADevice>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListOpenIDConnectProviderTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListOpenIDConnectProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderList` | `List<OpenIDConnectProviderListEntry>` | no |

## ListOrganizationsFeatures

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `EnabledFeatures` | `List<string>` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | no |
| `OnlyAttached` | `boolean` | no |
| `PathPrefix` | `string` | no |
| `PolicyUsageFilter` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `List<Policy>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListPoliciesGrantingServiceAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Arn` | `string` | yes |
| `ServiceNamespaces` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoliciesGrantingServiceAccess` | `List<ListPoliciesGrantingServiceAccessEntry>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListPolicyTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListPolicyVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Versions` | `List<PolicyVersion>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListRolePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyNames` | `List<string>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListRoleTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Roles` | `List<Role>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListSAMLProviderTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListSAMLProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderList` | `List<SAMLProviderListEntry>` | no |

## ListSSHPublicKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SSHPublicKeys` | `List<SSHPublicKeyMetadata>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListServerCertificateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListServerCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateMetadataList` | `List<ServerCertificateMetadata>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListServiceSpecificCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `ServiceName` | `string` | no |
| `AllUsers` | `boolean` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSpecificCredentials` | `List<ServiceSpecificCredentialMetadata>` | no |
| `Marker` | `string` | no |
| `IsTruncated` | `boolean` | no |

## ListSigningCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificates` | `List<SigningCertificate>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListUserPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyNames` | `List<string>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListUserTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathPrefix` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## ListVirtualMFADevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentStatus` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VirtualMFADevices` | `List<VirtualMFADevice>` | yes |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## PutAccountProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Properties` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRolePermissionsBoundary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PermissionsBoundary` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRolePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutUserPermissionsBoundary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PermissionsBoundary` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutUserPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RejectDelegationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequestId` | `string` | yes |
| `Notes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveClientIDFromOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |
| `ClientID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveRoleFromInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |
| `RoleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveUserFromGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetServiceSpecificCredential

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `ServiceSpecificCredentialId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSpecificCredential` | `ServiceSpecificCredential` | no |

## ResyncMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `SerialNumber` | `string` | yes |
| `AuthenticationCode1` | `string` | yes |
| `AuthenticationCode2` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendDelegationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetDefaultPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetSecurityTokenServicePreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalEndpointTokenVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SimulateCustomPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyInputList` | `List<string>` | yes |
| `PermissionsBoundaryPolicyInputList` | `List<string>` | no |
| `OrderedOrganizationPolicyInputList` | `List<OrderedOrganizationPolicyType>` | no |
| `ActionNames` | `List<string>` | yes |
| `ResourceArns` | `List<string>` | no |
| `ResourcePolicy` | `string` | no |
| `ResourceOwner` | `string` | no |
| `CallerArn` | `string` | no |
| `ContextEntries` | `List<ContextEntry>` | no |
| `ResourceHandlingOption` | `string` | no |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationResults` | `List<EvaluationResult>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## SimulatePrincipalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicySourceArn` | `string` | yes |
| `PolicyInputList` | `List<string>` | no |
| `PermissionsBoundaryPolicyInputList` | `List<string>` | no |
| `PolicyExclusionList` | `List<PolicyIdentifier>` | no |
| `ActionNames` | `List<string>` | yes |
| `ResourceArns` | `List<string>` | no |
| `ResourcePolicy` | `string` | no |
| `ResourceOwner` | `string` | no |
| `CallerArn` | `string` | no |
| `ContextEntries` | `List<ContextEntry>` | no |
| `ResourceHandlingOption` | `string` | no |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationResults` | `List<EvaluationResult>` | no |
| `IsTruncated` | `boolean` | no |
| `Marker` | `string` | no |

## TagInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialNumber` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagSAMLProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagServerCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagMFADevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialNumber` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagOpenIDConnectProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagSAMLProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagServerCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccessKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `AccessKeyId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccountPasswordPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MinimumPasswordLength` | `integer` | no |
| `RequireSymbols` | `boolean` | no |
| `RequireNumbers` | `boolean` | no |
| `RequireUppercaseCharacters` | `boolean` | no |
| `RequireLowercaseCharacters` | `boolean` | no |
| `AllowUsersToChangePassword` | `boolean` | no |
| `MaxPasswordAge` | `integer` | no |
| `PasswordReusePrevention` | `integer` | no |
| `HardExpiry` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAssumeRolePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `PolicyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDelegationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationRequestId` | `string` | yes |
| `Notes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `NewPath` | `string` | no |
| `NewGroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLoginProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `Password` | `string` | no |
| `PasswordResetRequired` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOpenIDConnectProviderThumbprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenIDConnectProviderArn` | `string` | yes |
| `ThumbprintList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `Description` | `string` | no |
| `MaxSessionDuration` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoleDescription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleName` | `string` | yes |
| `Description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `Role` | no |

## UpdateSAMLProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLMetadataDocument` | `string` | no |
| `SAMLProviderArn` | `string` | yes |
| `AssertionEncryptionMode` | `string` | no |
| `AddPrivateKey` | `string` | no |
| `RemovePrivateKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLProviderArn` | `string` | no |

## UpdateSSHPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `SSHPublicKeyId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServerCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateName` | `string` | yes |
| `NewPath` | `string` | no |
| `NewServerCertificateName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServiceSpecificCredential

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `ServiceSpecificCredentialId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSigningCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `CertificateId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `NewPath` | `string` | no |
| `NewUserName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UploadSSHPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `SSHPublicKeyBody` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SSHPublicKey` | `SSHPublicKey` | no |

## UploadServerCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | no |
| `ServerCertificateName` | `string` | yes |
| `CertificateBody` | `string` | yes |
| `PrivateKey` | `string` | yes |
| `CertificateChain` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerCertificateMetadata` | `ServerCertificateMetadata` | no |
| `Tags` | `List<Tag>` | no |

## UploadSigningCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `CertificateBody` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `SigningCertificate` | yes |

