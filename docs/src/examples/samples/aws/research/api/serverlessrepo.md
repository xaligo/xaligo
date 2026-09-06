# AWSServerlessApplicationRepository

API version: 2017-09-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/serverlessrepo/2017-09-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Author` | `string` | yes |
| `Description` | `string` | yes |
| `HomePageUrl` | `string` | no |
| `Labels` | `List<string>` | no |
| `LicenseBody` | `string` | no |
| `LicenseUrl` | `string` | no |
| `Name` | `string` | yes |
| `ReadmeBody` | `string` | no |
| `ReadmeUrl` | `string` | no |
| `SemanticVersion` | `string` | no |
| `SourceCodeArchiveUrl` | `string` | no |
| `SourceCodeUrl` | `string` | no |
| `SpdxLicenseId` | `string` | no |
| `TemplateBody` | `string` | no |
| `TemplateUrl` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Author` | `string` | no |
| `CreationTime` | `string` | no |
| `Description` | `string` | no |
| `HomePageUrl` | `string` | no |
| `IsVerifiedAuthor` | `boolean` | no |
| `Labels` | `List<string>` | no |
| `LicenseUrl` | `string` | no |
| `Name` | `string` | no |
| `ReadmeUrl` | `string` | no |
| `SpdxLicenseId` | `string` | no |
| `VerifiedAuthorUrl` | `string` | no |
| `Version` | `Version` | no |

## CreateApplicationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SemanticVersion` | `string` | yes |
| `SourceCodeArchiveUrl` | `string` | no |
| `SourceCodeUrl` | `string` | no |
| `TemplateBody` | `string` | no |
| `TemplateUrl` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `CreationTime` | `string` | no |
| `ParameterDefinitions` | `List<ParameterDefinition>` | no |
| `RequiredCapabilities` | `List<string>` | no |
| `ResourcesSupported` | `boolean` | no |
| `SemanticVersion` | `string` | no |
| `SourceCodeArchiveUrl` | `string` | no |
| `SourceCodeUrl` | `string` | no |
| `TemplateUrl` | `string` | no |

## CreateCloudFormationChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Capabilities` | `List<string>` | no |
| `ChangeSetName` | `string` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `NotificationArns` | `List<string>` | no |
| `ParameterOverrides` | `List<ParameterValue>` | no |
| `ResourceTypes` | `List<string>` | no |
| `RollbackConfiguration` | `RollbackConfiguration` | no |
| `SemanticVersion` | `string` | no |
| `StackName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `TemplateId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ChangeSetId` | `string` | no |
| `SemanticVersion` | `string` | no |
| `StackId` | `string` | no |

## CreateCloudFormationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SemanticVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `CreationTime` | `string` | no |
| `ExpirationTime` | `string` | no |
| `SemanticVersion` | `string` | no |
| `Status` | `string` | no |
| `TemplateId` | `string` | no |
| `TemplateUrl` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SemanticVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Author` | `string` | no |
| `CreationTime` | `string` | no |
| `Description` | `string` | no |
| `HomePageUrl` | `string` | no |
| `IsVerifiedAuthor` | `boolean` | no |
| `Labels` | `List<string>` | no |
| `LicenseUrl` | `string` | no |
| `Name` | `string` | no |
| `ReadmeUrl` | `string` | no |
| `SpdxLicenseId` | `string` | no |
| `VerifiedAuthorUrl` | `string` | no |
| `Version` | `Version` | no |

## GetApplicationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statements` | `List<ApplicationPolicyStatement>` | no |

## GetCloudFormationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `TemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `CreationTime` | `string` | no |
| `ExpirationTime` | `string` | no |
| `SemanticVersion` | `string` | no |
| `Status` | `string` | no |
| `TemplateId` | `string` | no |
| `TemplateUrl` | `string` | no |

## ListApplicationDependencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `MaxItems` | `integer` | no |
| `NextToken` | `string` | no |
| `SemanticVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Dependencies` | `List<ApplicationDependencySummary>` | no |
| `NextToken` | `string` | no |

## ListApplicationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `MaxItems` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionSummary>` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxItems` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<ApplicationSummary>` | no |
| `NextToken` | `string` | no |

## PutApplicationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Statements` | `List<ApplicationPolicyStatement>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statements` | `List<ApplicationPolicyStatement>` | no |

## UnshareApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Author` | `string` | no |
| `Description` | `string` | no |
| `HomePageUrl` | `string` | no |
| `Labels` | `List<string>` | no |
| `ReadmeBody` | `string` | no |
| `ReadmeUrl` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Author` | `string` | no |
| `CreationTime` | `string` | no |
| `Description` | `string` | no |
| `HomePageUrl` | `string` | no |
| `IsVerifiedAuthor` | `boolean` | no |
| `Labels` | `List<string>` | no |
| `LicenseUrl` | `string` | no |
| `Name` | `string` | no |
| `ReadmeUrl` | `string` | no |
| `SpdxLicenseId` | `string` | no |
| `VerifiedAuthorUrl` | `string` | no |
| `Version` | `Version` | no |

