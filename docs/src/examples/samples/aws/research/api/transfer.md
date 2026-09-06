# AWS Transfer Family

API version: 2018-11-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/transfer/2018-11-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeDirectory` | `string` | no |
| `HomeDirectoryType` | `string` | no |
| `HomeDirectoryMappings` | `List<HomeDirectoryMapEntry>` | no |
| `Policy` | `string` | no |
| `PosixProfile` | `PosixProfile` | no |
| `Role` | `string` | yes |
| `ServerId` | `string` | yes |
| `ExternalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `ExternalId` | `string` | yes |

## CreateAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `ServerId` | `string` | yes |
| `LocalProfileId` | `string` | yes |
| `PartnerProfileId` | `string` | yes |
| `BaseDirectory` | `string` | no |
| `AccessRole` | `string` | yes |
| `Status` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `PreserveFilename` | `string` | no |
| `EnforceMessageSigning` | `string` | no |
| `CustomDirectories` | `CustomDirectoriesType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgreementId` | `string` | yes |

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |
| `As2Config` | `As2ConnectorConfig` | no |
| `AccessRole` | `string` | yes |
| `LoggingRole` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SftpConfig` | `SftpConnectorConfig` | no |
| `SecurityPolicyName` | `string` | no |
| `EgressConfig` | `ConnectorEgressConfig` | no |
| `IpAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

## CreateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `As2Id` | `string` | yes |
| `ProfileType` | `string` | yes |
| `CertificateIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

## CreateServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |
| `Domain` | `string` | no |
| `EndpointDetails` | `EndpointDetails` | no |
| `EndpointType` | `string` | no |
| `HostKey` | `string` | no |
| `IdentityProviderDetails` | `IdentityProviderDetails` | no |
| `IdentityProviderType` | `string` | no |
| `LoggingRole` | `string` | no |
| `PostAuthenticationLoginBanner` | `string` | no |
| `PreAuthenticationLoginBanner` | `string` | no |
| `Protocols` | `List<string>` | no |
| `ProtocolDetails` | `ProtocolDetails` | no |
| `SecurityPolicyName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `WorkflowDetails` | `WorkflowDetails` | no |
| `StructuredLogDestinations` | `List<string>` | no |
| `S3StorageOptions` | `S3StorageOptions` | no |
| `IpAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeDirectory` | `string` | no |
| `HomeDirectoryType` | `string` | no |
| `HomeDirectoryMappings` | `List<HomeDirectoryMapEntry>` | no |
| `Policy` | `string` | no |
| `PosixProfile` | `PosixProfile` | no |
| `Role` | `string` | yes |
| `ServerId` | `string` | yes |
| `SshPublicKeyBody` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `UserName` | `string` | yes |

## CreateWebApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderDetails` | `WebAppIdentityProviderDetails` | yes |
| `AccessEndpoint` | `string` | no |
| `WebAppUnits` | `WebAppUnits` | no |
| `Tags` | `List<Tag>` | no |
| `WebAppEndpointPolicy` | `string` | no |
| `EndpointDetails` | `WebAppEndpointDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

## CreateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Steps` | `List<WorkflowStep>` | yes |
| `OnExceptionSteps` | `List<WorkflowStep>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | yes |

## DeleteAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `ExternalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgreementId` | `string` | yes |
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHostKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `HostKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSshPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `SshPublicKeyId` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWebApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWebAppCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `ExternalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `Access` | `DescribedAccess` | yes |

## DescribeAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgreementId` | `string` | yes |
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Agreement` | `DescribedAgreement` | yes |

## DescribeCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `DescribedCertificate` | yes |

## DescribeConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connector` | `DescribedConnector` | yes |

## DescribeExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionId` | `string` | yes |
| `WorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | yes |
| `Execution` | `DescribedExecution` | yes |

## DescribeHostKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `HostKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostKey` | `DescribedHostKey` | yes |

## DescribeProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profile` | `DescribedProfile` | yes |

## DescribeSecurityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityPolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityPolicy` | `DescribedSecurityPolicy` | yes |

## DescribeServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Server` | `DescribedServer` | yes |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `User` | `DescribedUser` | yes |

## DescribeWebApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebApp` | `DescribedWebApp` | yes |

## DescribeWebAppCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppCustomization` | `DescribedWebAppCustomization` | yes |

## DescribeWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workflow` | `DescribedWorkflow` | yes |

## ImportCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Usage` | `string` | yes |
| `Certificate` | `string` | yes |
| `CertificateChain` | `string` | no |
| `PrivateKey` | `string` | no |
| `ActiveDate` | `timestamp` | no |
| `InactiveDate` | `timestamp` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateId` | `string` | yes |

## ImportHostKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `HostKeyBody` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `HostKeyId` | `string` | yes |

## ImportSshPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `SshPublicKeyBody` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `SshPublicKeyId` | `string` | yes |
| `UserName` | `string` | yes |

## ListAccesses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |
| `Accesses` | `List<ListedAccess>` | yes |

## ListAgreements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Agreements` | `List<ListedAgreement>` | yes |

## ListCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Certificates` | `List<ListedCertificate>` | yes |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Connectors` | `List<ListedConnector>` | yes |

## ListExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `WorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WorkflowId` | `string` | yes |
| `Executions` | `List<ListedExecution>` | yes |

## ListFileTransferResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `TransferId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileTransferResults` | `List<ConnectorFileTransferResult>` | yes |
| `NextToken` | `string` | no |

## ListHostKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |
| `HostKeys` | `List<ListedHostKey>` | yes |

## ListProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ProfileType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Profiles` | `List<ListedProfile>` | yes |

## ListSecurityPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SecurityPolicyNames` | `List<string>` | yes |

## ListServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Servers` | `List<ListedServer>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `NextToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServerId` | `string` | yes |
| `Users` | `List<ListedUser>` | yes |

## ListWebApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WebApps` | `List<ListedWebApp>` | yes |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Workflows` | `List<ListedWorkflow>` | yes |

## SendWorkflowStepState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | yes |
| `ExecutionId` | `string` | yes |
| `Token` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartDirectoryListing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `RemoteDirectoryPath` | `string` | yes |
| `MaxItems` | `integer` | no |
| `OutputDirectoryPath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListingId` | `string` | yes |
| `OutputFileName` | `string` | yes |

## StartFileTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `SendFilePaths` | `List<string>` | no |
| `RetrieveFilePaths` | `List<string>` | no |
| `LocalDirectoryPath` | `string` | no |
| `RemoteDirectoryPath` | `string` | no |
| `CustomHttpHeaders` | `List<CustomHttpHeader>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransferId` | `string` | yes |

## StartRemoteDelete

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `DeletePath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeleteId` | `string` | yes |

## StartRemoteMove

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `SourcePath` | `string` | yes |
| `TargetPath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MoveId` | `string` | yes |

## StartServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `SftpConnectionDetails` | `SftpConnectorConnectionDetails` | no |

## TestIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `ServerProtocol` | `string` | no |
| `SourceIp` | `string` | no |
| `UserName` | `string` | yes |
| `UserPassword` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Response` | `string` | no |
| `StatusCode` | `integer` | yes |
| `Message` | `string` | no |
| `Url` | `string` | yes |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeDirectory` | `string` | no |
| `HomeDirectoryType` | `string` | no |
| `HomeDirectoryMappings` | `List<HomeDirectoryMapEntry>` | no |
| `Policy` | `string` | no |
| `PosixProfile` | `PosixProfile` | no |
| `Role` | `string` | no |
| `ServerId` | `string` | yes |
| `ExternalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `ExternalId` | `string` | yes |

## UpdateAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgreementId` | `string` | yes |
| `ServerId` | `string` | yes |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `LocalProfileId` | `string` | no |
| `PartnerProfileId` | `string` | no |
| `BaseDirectory` | `string` | no |
| `AccessRole` | `string` | no |
| `PreserveFilename` | `string` | no |
| `EnforceMessageSigning` | `string` | no |
| `CustomDirectories` | `CustomDirectoriesType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgreementId` | `string` | yes |

## UpdateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateId` | `string` | yes |
| `ActiveDate` | `timestamp` | no |
| `InactiveDate` | `timestamp` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateId` | `string` | yes |

## UpdateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `Url` | `string` | no |
| `As2Config` | `As2ConnectorConfig` | no |
| `AccessRole` | `string` | no |
| `LoggingRole` | `string` | no |
| `SftpConfig` | `SftpConnectorConfig` | no |
| `SecurityPolicyName` | `string` | no |
| `EgressConfig` | `UpdateConnectorEgressConfig` | no |
| `IpAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

## UpdateHostKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `HostKeyId` | `string` | yes |
| `Description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `HostKeyId` | `string` | yes |

## UpdateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `CertificateIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

## UpdateServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |
| `ProtocolDetails` | `ProtocolDetails` | no |
| `EndpointDetails` | `EndpointDetails` | no |
| `EndpointType` | `string` | no |
| `HostKey` | `string` | no |
| `IdentityProviderDetails` | `IdentityProviderDetails` | no |
| `LoggingRole` | `string` | no |
| `PostAuthenticationLoginBanner` | `string` | no |
| `PreAuthenticationLoginBanner` | `string` | no |
| `Protocols` | `List<string>` | no |
| `SecurityPolicyName` | `string` | no |
| `ServerId` | `string` | yes |
| `WorkflowDetails` | `WorkflowDetails` | no |
| `StructuredLogDestinations` | `List<string>` | no |
| `S3StorageOptions` | `S3StorageOptions` | no |
| `IpAddressType` | `string` | no |
| `IdentityProviderType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeDirectory` | `string` | no |
| `HomeDirectoryType` | `string` | no |
| `HomeDirectoryMappings` | `List<HomeDirectoryMapEntry>` | no |
| `Policy` | `string` | no |
| `PosixProfile` | `PosixProfile` | no |
| `Role` | `string` | no |
| `ServerId` | `string` | yes |
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerId` | `string` | yes |
| `UserName` | `string` | yes |

## UpdateWebApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |
| `IdentityProviderDetails` | `UpdateWebAppIdentityProviderDetails` | no |
| `AccessEndpoint` | `string` | no |
| `WebAppUnits` | `WebAppUnits` | no |
| `EndpointDetails` | `UpdateWebAppEndpointDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

## UpdateWebAppCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |
| `Title` | `string` | no |
| `LogoFile` | `blob` | no |
| `FaviconFile` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAppId` | `string` | yes |

