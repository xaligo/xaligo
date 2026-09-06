# Amazon WorkMail

API version: 2017-10-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workmail/2017-10-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateDelegateToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `EntityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateMemberToGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `GroupId` | `string` | yes |
| `MemberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssumeImpersonationRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ImpersonationRoleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Token` | `string` | no |
| `ExpiresIn` | `long` | no |

## CancelMailboxExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `JobId` | `string` | yes |
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `Alias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAvailabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |
| `EwsProvider` | `EwsAvailabilityProvider` | no |
| `LambdaProvider` | `LambdaAvailabilityProvider` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Name` | `string` | yes |
| `HiddenFromGlobalAddressList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | no |

## CreateIdentityCenterApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `InstanceArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | no |

## CreateImpersonationRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `OrganizationId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `List<ImpersonationRule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImpersonationRoleId` | `string` | no |

## CreateMobileDeviceAccessRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Effect` | `string` | yes |
| `DeviceTypes` | `List<string>` | no |
| `NotDeviceTypes` | `List<string>` | no |
| `DeviceModels` | `List<string>` | no |
| `NotDeviceModels` | `List<string>` | no |
| `DeviceOperatingSystems` | `List<string>` | no |
| `NotDeviceOperatingSystems` | `List<string>` | no |
| `DeviceUserAgents` | `List<string>` | no |
| `NotDeviceUserAgents` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MobileDeviceAccessRuleId` | `string` | no |

## CreateOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `Alias` | `string` | yes |
| `ClientToken` | `string` | no |
| `Domains` | `List<Domain>` | no |
| `KmsKeyArn` | `string` | no |
| `EnableInteroperability` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |

## CreateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Description` | `string` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Name` | `string` | yes |
| `DisplayName` | `string` | yes |
| `Password` | `string` | no |
| `Role` | `string` | no |
| `FirstName` | `string` | no |
| `LastName` | `string` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |
| `IdentityProviderUserId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |

## DeleteAccessControlRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `Alias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAvailabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEmailMonitoringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdentityCenterApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdentityProviderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImpersonationRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ImpersonationRoleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMailboxPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `GranteeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMobileDeviceAccessOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |
| `DeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMobileDeviceAccessRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `MobileDeviceAccessRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `OrganizationId` | `string` | yes |
| `DeleteDirectory` | `boolean` | yes |
| `ForceDelete` | `boolean` | no |
| `DeleteIdentityCenterApplication` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `State` | `string` | no |

## DeletePersonalAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `PersonalAccessTokenId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRetentionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterFromWorkMail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterMailDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeEmailMonitoringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | no |
| `LogGroupArn` | `string` | no |

## DescribeEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Email` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityId` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |

## DescribeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | no |
| `Name` | `string` | no |
| `Email` | `string` | no |
| `State` | `string` | no |
| `EnabledDate` | `timestamp` | no |
| `DisabledDate` | `timestamp` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |

## DescribeIdentityProviderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationMode` | `string` | no |
| `IdentityCenterConfiguration` | `IdentityCenterConfiguration` | no |
| `PersonalAccessTokenConfiguration` | `PersonalAccessTokenConfiguration` | no |

## DescribeInboundDmarcSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Enforced` | `boolean` | no |

## DescribeMailboxExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityId` | `string` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `S3BucketName` | `string` | no |
| `S3Prefix` | `string` | no |
| `S3Path` | `string` | no |
| `EstimatedProgress` | `integer` | no |
| `State` | `string` | no |
| `ErrorInfo` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |

## DescribeOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `Alias` | `string` | no |
| `State` | `string` | no |
| `DirectoryId` | `string` | no |
| `DirectoryType` | `string` | no |
| `DefaultMailDomain` | `string` | no |
| `CompletedDate` | `timestamp` | no |
| `ErrorMessage` | `string` | no |
| `ARN` | `string` | no |
| `MigrationAdmin` | `string` | no |
| `InteroperabilityEnabled` | `boolean` | no |

## DescribeResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |
| `Email` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `BookingOptions` | `BookingOptions` | no |
| `State` | `string` | no |
| `EnabledDate` | `timestamp` | no |
| `DisabledDate` | `timestamp` | no |
| `Description` | `string` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `Name` | `string` | no |
| `Email` | `string` | no |
| `DisplayName` | `string` | no |
| `State` | `string` | no |
| `UserRole` | `string` | no |
| `EnabledDate` | `timestamp` | no |
| `DisabledDate` | `timestamp` | no |
| `MailboxProvisionedDate` | `timestamp` | no |
| `MailboxDeprovisionedDate` | `timestamp` | no |
| `FirstName` | `string` | no |
| `LastName` | `string` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |
| `Initials` | `string` | no |
| `Telephone` | `string` | no |
| `Street` | `string` | no |
| `JobTitle` | `string` | no |
| `City` | `string` | no |
| `Company` | `string` | no |
| `ZipCode` | `string` | no |
| `Department` | `string` | no |
| `Country` | `string` | no |
| `Office` | `string` | no |
| `IdentityProviderUserId` | `string` | no |
| `IdentityProviderIdentityStoreId` | `string` | no |

## DisassociateDelegateFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `EntityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMemberFromGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `GroupId` | `string` | yes |
| `MemberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccessControlEffect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `IpAddress` | `string` | yes |
| `Action` | `string` | yes |
| `UserId` | `string` | no |
| `ImpersonationRoleId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Effect` | `string` | no |
| `MatchedRules` | `List<string>` | no |

## GetDefaultRetentionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `FolderConfigurations` | `List<FolderConfiguration>` | no |

## GetImpersonationRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ImpersonationRoleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImpersonationRoleId` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `Description` | `string` | no |
| `Rules` | `List<ImpersonationRule>` | no |
| `DateCreated` | `timestamp` | no |
| `DateModified` | `timestamp` | no |

## GetImpersonationRoleEffect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ImpersonationRoleId` | `string` | yes |
| `TargetUser` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Effect` | `string` | no |
| `MatchedRules` | `List<ImpersonationMatchedRule>` | no |

## GetMailDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<DnsRecord>` | no |
| `IsTestDomain` | `boolean` | no |
| `IsDefault` | `boolean` | no |
| `OwnershipVerificationStatus` | `string` | no |
| `DkimVerificationStatus` | `string` | no |

## GetMailboxDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MailboxQuota` | `integer` | no |
| `MailboxSize` | `double` | no |

## GetMobileDeviceAccessEffect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DeviceType` | `string` | no |
| `DeviceModel` | `string` | no |
| `DeviceOperatingSystem` | `string` | no |
| `DeviceUserAgent` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Effect` | `string` | no |
| `MatchedRules` | `List<MobileDeviceAccessMatchedRule>` | no |

## GetMobileDeviceAccessOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |
| `DeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `DeviceId` | `string` | no |
| `Effect` | `string` | no |
| `Description` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateModified` | `timestamp` | no |

## GetPersonalAccessTokenMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `PersonalAccessTokenId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersonalAccessTokenId` | `string` | no |
| `UserId` | `string` | no |
| `Name` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateLastUsed` | `timestamp` | no |
| `ExpiresTime` | `timestamp` | no |
| `Scopes` | `List<string>` | no |

## ListAccessControlRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<AccessControlRule>` | no |

## ListAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListAvailabilityConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityConfigurations` | `List<AvailabilityConfiguration>` | no |
| `NextToken` | `string` | no |

## ListGroupMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `GroupId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Members` | `List<Member>` | no |
| `NextToken` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `ListGroupsFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<Group>` | no |
| `NextToken` | `string` | no |

## ListGroupsForEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `Filters` | `ListGroupsForEntityFilters` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupIdentifier>` | no |
| `NextToken` | `string` | no |

## ListImpersonationRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Roles` | `List<ImpersonationRole>` | no |
| `NextToken` | `string` | no |

## ListMailDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MailDomains` | `List<MailDomainSummary>` | no |
| `NextToken` | `string` | no |

## ListMailboxExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<MailboxExportJob>` | no |
| `NextToken` | `string` | no |

## ListMailboxPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Permissions` | `List<Permission>` | no |
| `NextToken` | `string` | no |

## ListMobileDeviceAccessOverrides

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | no |
| `DeviceId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Overrides` | `List<MobileDeviceAccessOverride>` | no |
| `NextToken` | `string` | no |

## ListMobileDeviceAccessRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<MobileDeviceAccessRule>` | no |

## ListOrganizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationSummaries` | `List<OrganizationSummary>` | no |
| `NextToken` | `string` | no |

## ListPersonalAccessTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PersonalAccessTokenSummaries` | `List<PersonalAccessTokenSummary>` | no |

## ListResourceDelegates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Delegates` | `List<Delegate>` | no |
| `NextToken` | `string` | no |

## ListResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `ListResourcesFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<Resource>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `ListUsersFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `NextToken` | `string` | no |

## PutAccessControlRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Effect` | `string` | yes |
| `Description` | `string` | yes |
| `IpRanges` | `List<string>` | no |
| `NotIpRanges` | `List<string>` | no |
| `Actions` | `List<string>` | no |
| `NotActions` | `List<string>` | no |
| `UserIds` | `List<string>` | no |
| `NotUserIds` | `List<string>` | no |
| `OrganizationId` | `string` | yes |
| `ImpersonationRoleIds` | `List<string>` | no |
| `NotImpersonationRoleIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEmailMonitoringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `RoleArn` | `string` | no |
| `LogGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutIdentityProviderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `AuthenticationMode` | `string` | yes |
| `IdentityCenterConfiguration` | `IdentityCenterConfiguration` | yes |
| `PersonalAccessTokenConfiguration` | `PersonalAccessTokenConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutInboundDmarcSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Enforced` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMailboxPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `GranteeId` | `string` | yes |
| `PermissionValues` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMobileDeviceAccessOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `Effect` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRetentionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Id` | `string` | no |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `FolderConfigurations` | `List<FolderConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterMailDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterToWorkMail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `Email` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |
| `Password` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartMailboxExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `Description` | `string` | no |
| `RoleArn` | `string` | yes |
| `KmsKeyArn` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `S3Prefix` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestAvailabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | no |
| `EwsProvider` | `EwsAvailabilityProvider` | no |
| `LambdaProvider` | `LambdaAvailabilityProvider` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestPassed` | `boolean` | no |
| `FailureReason` | `string` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAvailabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |
| `EwsProvider` | `EwsAvailabilityProvider` | no |
| `LambdaProvider` | `LambdaAvailabilityProvider` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDefaultMailDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `GroupId` | `string` | yes |
| `HiddenFromGlobalAddressList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateImpersonationRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ImpersonationRoleId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `List<ImpersonationRule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMailboxQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |
| `MailboxQuota` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMobileDeviceAccessRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `MobileDeviceAccessRuleId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Effect` | `string` | yes |
| `DeviceTypes` | `List<string>` | no |
| `NotDeviceTypes` | `List<string>` | no |
| `DeviceModels` | `List<string>` | no |
| `NotDeviceModels` | `List<string>` | no |
| `DeviceOperatingSystems` | `List<string>` | no |
| `NotDeviceOperatingSystems` | `List<string>` | no |
| `DeviceUserAgents` | `List<string>` | no |
| `NotDeviceUserAgents` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePrimaryEmailAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `EntityId` | `string` | yes |
| `Email` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `Name` | `string` | no |
| `BookingOptions` | `BookingOptions` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `UserId` | `string` | yes |
| `Role` | `string` | no |
| `DisplayName` | `string` | no |
| `FirstName` | `string` | no |
| `LastName` | `string` | no |
| `HiddenFromGlobalAddressList` | `boolean` | no |
| `Initials` | `string` | no |
| `Telephone` | `string` | no |
| `Street` | `string` | no |
| `JobTitle` | `string` | no |
| `City` | `string` | no |
| `Company` | `string` | no |
| `ZipCode` | `string` | no |
| `Department` | `string` | no |
| `Country` | `string` | no |
| `Office` | `string` | no |
| `IdentityProviderUserId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


