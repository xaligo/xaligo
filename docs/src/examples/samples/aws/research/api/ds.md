# AWS Directory Service

API version: 2015-04-16. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ds/2015-04-16/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptSharedDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectory` | `SharedDirectory` | no |

## AddIpRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `IpRoutes` | `List<IpRoute>` | yes |
| `UpdateSecurityGroupForDirectoryControllers` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RegionName` | `string` | yes |
| `VPCSettings` | `DirectoryVpcSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelSchemaExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `SchemaExtensionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ConnectDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ShortName` | `string` | no |
| `Password` | `string` | yes |
| `Description` | `string` | no |
| `Size` | `string` | yes |
| `ConnectSettings` | `DirectoryConnectSettings` | yes |
| `Tags` | `List<Tag>` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |

## CreateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Alias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `Alias` | `string` | no |

## CreateComputer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `ComputerName` | `string` | yes |
| `Password` | `string` | yes |
| `OrganizationalUnitDistinguishedName` | `string` | no |
| `ComputerAttributes` | `List<Attribute>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Computer` | `Computer` | no |

## CreateConditionalForwarder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RemoteDomainName` | `string` | yes |
| `DnsIpAddrs` | `List<string>` | no |
| `DnsIpv6Addrs` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ShortName` | `string` | no |
| `Password` | `string` | yes |
| `Description` | `string` | no |
| `Size` | `string` | yes |
| `VpcSettings` | `DirectoryVpcSettings` | no |
| `Tags` | `List<Tag>` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |

## CreateHybridAD

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretArn` | `string` | yes |
| `AssessmentId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |

## CreateLogSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `LogGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateMicrosoftAD

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ShortName` | `string` | no |
| `Password` | `string` | yes |
| `Description` | `string` | no |
| `VpcSettings` | `DirectoryVpcSettings` | yes |
| `Edition` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |

## CreateTrust

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RemoteDomainName` | `string` | yes |
| `TrustPassword` | `string` | yes |
| `TrustDirection` | `string` | yes |
| `TrustType` | `string` | no |
| `ConditionalForwarderIpAddrs` | `List<string>` | no |
| `ConditionalForwarderIpv6Addrs` | `List<string>` | no |
| `SelectiveAuth` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustId` | `string` | no |

## DeleteADAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentId` | `string` | no |

## DeleteConditionalForwarder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RemoteDomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |

## DeleteLogSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |

## DeleteTrust

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustId` | `string` | yes |
| `DeleteAssociatedConditionalForwarder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustId` | `string` | no |

## DeregisterCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `CertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterEventTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `TopicName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeADAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assessment` | `Assessment` | no |
| `AssessmentReports` | `List<AssessmentReport>` | no |

## DescribeCAEnrollmentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `PcaConnectorArn` | `string` | no |
| `CaEnrollmentPolicyStatus` | `string` | no |
| `LastUpdatedDateTime` | `timestamp` | no |
| `CaEnrollmentPolicyStatusReason` | `string` | no |

## DescribeCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `CertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `Certificate` | no |

## DescribeClientAuthenticationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Type` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientAuthenticationSettingsInfo` | `List<ClientAuthenticationSettingInfo>` | no |
| `NextToken` | `string` | no |

## DescribeConditionalForwarders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RemoteDomainNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConditionalForwarders` | `List<ConditionalForwarder>` | no |

## DescribeDirectories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryDescriptions` | `List<DirectoryDescription>` | no |
| `NextToken` | `string` | no |

## DescribeDirectoryDataAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataAccessStatus` | `string` | no |

## DescribeDomainControllers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `DomainControllerIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainControllers` | `List<DomainController>` | no |
| `NextToken` | `string` | no |

## DescribeEventTopics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `TopicNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventTopics` | `List<EventTopic>` | no |

## DescribeHybridADUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UpdateType` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateActivities` | `HybridUpdateActivities` | no |
| `NextToken` | `string` | no |

## DescribeLDAPSSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Type` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LDAPSSettingsInfo` | `List<LDAPSSettingInfo>` | no |
| `NextToken` | `string` | no |

## DescribeRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RegionName` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionsDescription` | `List<RegionDescription>` | no |
| `NextToken` | `string` | no |

## DescribeSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Status` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `SettingEntries` | `List<SettingEntry>` | no |
| `NextToken` | `string` | no |

## DescribeSharedDirectories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerDirectoryId` | `string` | yes |
| `SharedDirectoryIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectories` | `List<SharedDirectory>` | no |
| `NextToken` | `string` | no |

## DescribeSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `SnapshotIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshots` | `List<Snapshot>` | no |
| `NextToken` | `string` | no |

## DescribeTrusts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `TrustIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Trusts` | `List<Trust>` | no |
| `NextToken` | `string` | no |

## DescribeUpdateDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UpdateType` | `string` | yes |
| `RegionName` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateActivities` | `List<UpdateInfoEntry>` | no |
| `NextToken` | `string` | no |

## DisableCAEnrollmentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableClientAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableDirectoryDataAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableLDAPS

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableRadius

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableSso

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UserName` | `string` | no |
| `Password` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableCAEnrollmentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `PcaConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableClientAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableDirectoryDataAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableLDAPS

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableRadius

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RadiusSettings` | `RadiusSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableSso

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UserName` | `string` | no |
| `Password` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDirectoryLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryLimits` | `DirectoryLimits` | no |

## GetSnapshotLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotLimits` | `SnapshotLimits` | no |

## ListADAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assessments` | `List<AssessmentSummary>` | no |
| `NextToken` | `string` | no |

## ListCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CertificatesInfo` | `List<CertificateInfo>` | no |

## ListIpRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpRoutesInfo` | `List<IpRouteInfo>` | no |
| `NextToken` | `string` | no |

## ListLogSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogSubscriptions` | `List<LogSubscription>` | no |
| `NextToken` | `string` | no |

## ListSchemaExtensions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaExtensionsInfo` | `List<SchemaExtensionInfo>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## RegisterCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `CertificateData` | `string` | yes |
| `Type` | `string` | no |
| `ClientCertAuthSettings` | `ClientCertAuthSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateId` | `string` | no |

## RegisterEventTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `TopicName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RejectSharedDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectoryId` | `string` | no |

## RemoveIpRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `CidrIps` | `List<string>` | no |
| `CidrIpv6s` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetUserPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UserName` | `string` | yes |
| `NewPassword` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestoreFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ShareDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `ShareNotes` | `string` | no |
| `ShareTarget` | `ShareTarget` | yes |
| `ShareMethod` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectoryId` | `string` | no |

## StartADAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentConfiguration` | `AssessmentConfiguration` | no |
| `DirectoryId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentId` | `string` | no |

## StartSchemaExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `CreateSnapshotBeforeSchemaExtension` | `boolean` | yes |
| `LdifContent` | `string` | yes |
| `Description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaExtensionId` | `string` | no |

## UnshareDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UnshareTarget` | `UnshareTarget` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SharedDirectoryId` | `string` | no |

## UpdateConditionalForwarder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RemoteDomainName` | `string` | yes |
| `DnsIpAddrs` | `List<string>` | no |
| `DnsIpv6Addrs` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDirectorySetup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `UpdateType` | `string` | yes |
| `OSUpdateSettings` | `OSUpdateSettings` | no |
| `DirectorySizeUpdateSettings` | `DirectorySizeUpdateSettings` | no |
| `NetworkUpdateSettings` | `NetworkUpdateSettings` | no |
| `CreateSnapshotBeforeUpdate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHybridAD

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `HybridAdministratorAccountUpdate` | `HybridAdministratorAccountUpdate` | no |
| `SelfManagedInstancesSettings` | `HybridCustomerInstancesSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `AssessmentId` | `string` | no |

## UpdateNumberOfDomainControllers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `DesiredNumber` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRadius

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `RadiusSettings` | `RadiusSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `Settings` | `List<Setting>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |

## UpdateTrust

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustId` | `string` | yes |
| `SelectiveAuth` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `TrustId` | `string` | no |

## VerifyTrust

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustId` | `string` | no |

