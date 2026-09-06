# Partner Central Account API

API version: 2025-04-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/partnercentral-account/2025-04-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptConnectionInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | yes |

## AssociateAwsTrainingCertificationEmailDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `Email` | `string` | yes |
| `EmailVerificationCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ConnectionType` | `string` | yes |
| `Reason` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `OtherParticipantAccountId` | `string` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `ConnectionTypes` | `Map<ConnectionTypeDetail>` | yes |

## CancelConnectionInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `ExpiresAt` | `timestamp` | no |
| `OtherParticipantIdentifier` | `string` | yes |
| `ParticipantType` | `string` | yes |
| `Status` | `string` | yes |
| `InvitationMessage` | `string` | yes |
| `InviterEmail` | `string` | yes |
| `InviterName` | `string` | yes |

## CancelProfileUpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `TaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `TaskDetails` | `TaskDetails` | yes |
| `StartedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `EndedAt` | `timestamp` | no |
| `ErrorDetailList` | `List<ErrorDetail>` | no |

## CreateConnectionInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `ConnectionType` | `string` | yes |
| `Email` | `string` | yes |
| `Message` | `string` | yes |
| `Name` | `string` | yes |
| `ReceiverIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `ExpiresAt` | `timestamp` | no |
| `OtherParticipantIdentifier` | `string` | yes |
| `ParticipantType` | `string` | yes |
| `Status` | `string` | yes |
| `InvitationMessage` | `string` | yes |
| `InviterEmail` | `string` | yes |
| `InviterName` | `string` | yes |

## CreatePartner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | no |
| `LegalName` | `string` | yes |
| `PrimarySolutionType` | `string` | yes |
| `AllianceLeadContact` | `AllianceLeadContact` | yes |
| `EmailVerificationCode` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `LegalName` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `Profile` | `PartnerProfile` | yes |
| `AwsTrainingCertificationEmailDomains` | `List<PartnerDomain>` | no |
| `AllianceLeadContact` | `AllianceLeadContact` | yes |

## DisassociateAwsTrainingCertificationEmailDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAllianceLeadContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `AllianceLeadContact` | `AllianceLeadContact` | yes |

## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `OtherParticipantAccountId` | `string` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `ConnectionTypes` | `Map<ConnectionTypeDetail>` | yes |

## GetConnectionInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `ExpiresAt` | `timestamp` | no |
| `OtherParticipantIdentifier` | `string` | yes |
| `ParticipantType` | `string` | yes |
| `Status` | `string` | yes |
| `InvitationMessage` | `string` | yes |
| `InviterEmail` | `string` | yes |
| `InviterName` | `string` | yes |

## GetConnectionPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `AccessType` | `string` | yes |
| `ExcludedParticipantIds` | `List<string>` | no |
| `UpdatedAt` | `timestamp` | yes |
| `Revision` | `long` | yes |

## GetPartner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `LegalName` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `Profile` | `PartnerProfile` | yes |
| `AwsTrainingCertificationEmailDomains` | `List<PartnerDomain>` | no |

## GetProfileUpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `TaskDetails` | `TaskDetails` | yes |
| `StartedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `EndedAt` | `timestamp` | no |
| `ErrorDetailList` | `List<ErrorDetail>` | no |

## GetProfileVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `Visibility` | `string` | yes |
| `ProfileId` | `string` | yes |

## GetQualificationsAssociationDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `Status` | `string` | yes |
| `PrimaryPartner` | `QualificationsAssociationPartner` | no |
| `AssociatedPartners` | `List<QualificationsAssociationPartner>` | no |
| `UpdatedAt` | `timestamp` | no |

## GetQualificationsAssociationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `Status` | `string` | yes |
| `PrimaryPartner` | `QualificationsAssociationPartner` | yes |
| `StartedAt` | `timestamp` | yes |
| `EndedAt` | `timestamp` | no |

## GetQualificationsDisassociationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `Status` | `string` | yes |
| `AssociatedPartner` | `QualificationsAssociationPartner` | yes |
| `StartedAt` | `timestamp` | yes |
| `EndedAt` | `timestamp` | no |

## GetVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationType` | `string` | yes |
| `VerificationStatus` | `string` | yes |
| `VerificationStatusReason` | `string` | no |
| `VerificationResponseDetails` | `VerificationResponseDetails` | yes |
| `StartedAt` | `timestamp` | yes |
| `CompletedAt` | `timestamp` | no |

## ListConnectionInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `NextToken` | `string` | no |
| `ConnectionType` | `string` | no |
| `MaxResults` | `integer` | no |
| `OtherParticipantIdentifiers` | `List<string>` | no |
| `ParticipantType` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionInvitationSummaries` | `List<ConnectionInvitationSummary>` | yes |
| `NextToken` | `string` | no |

## ListConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `NextToken` | `string` | no |
| `ConnectionType` | `string` | no |
| `MaxResults` | `integer` | no |
| `OtherParticipantIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionSummaries` | `List<ConnectionSummary>` | yes |
| `NextToken` | `string` | no |

## ListPartners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartnerSummaryList` | `List<PartnerSummary>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

## PutAllianceLeadContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `AllianceLeadContact` | `AllianceLeadContact` | yes |
| `EmailVerificationCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `AllianceLeadContact` | `AllianceLeadContact` | yes |

## PutProfileVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `Visibility` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `Visibility` | `string` | yes |
| `ProfileId` | `string` | yes |

## RejectConnectionInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `ExpiresAt` | `timestamp` | no |
| `OtherParticipantIdentifier` | `string` | yes |
| `ParticipantType` | `string` | yes |
| `Status` | `string` | yes |
| `InvitationMessage` | `string` | yes |
| `InviterEmail` | `string` | yes |
| `InviterName` | `string` | yes |

## SendEmailVerificationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Email` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartProfileUpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `TaskDetails` | `TaskDetails` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `TaskDetails` | `TaskDetails` | yes |
| `StartedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `EndedAt` | `timestamp` | no |
| `ErrorDetailList` | `List<ErrorDetail>` | no |

## StartQualificationsAssociationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `PrimaryPartner` | `QualificationsAssociationPartner` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `Status` | `string` | yes |
| `PrimaryPartner` | `QualificationsAssociationPartner` | yes |
| `StartedAt` | `timestamp` | yes |

## StartQualificationsDisassociationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `AssociatedPartner` | `QualificationsAssociationPartner` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `TaskId` | `string` | yes |
| `Status` | `string` | yes |
| `AssociatedPartner` | `QualificationsAssociationPartner` | yes |
| `StartedAt` | `timestamp` | yes |

## StartVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `VerificationDetails` | `VerificationDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationType` | `string` | yes |
| `VerificationStatus` | `string` | yes |
| `VerificationStatusReason` | `string` | no |
| `VerificationResponseDetails` | `VerificationResponseDetails` | yes |
| `StartedAt` | `timestamp` | yes |
| `CompletedAt` | `timestamp` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateConnectionPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Revision` | `long` | yes |
| `AccessType` | `string` | yes |
| `ExcludedParticipantIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Arn` | `string` | yes |
| `AccessType` | `string` | yes |
| `ExcludedParticipantIds` | `List<string>` | no |
| `UpdatedAt` | `timestamp` | yes |
| `Revision` | `long` | yes |

