# AWS Secrets Manager

API version: 2017-10-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/secretsmanager/2017-10-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetSecretValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretIdList` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretValues` | `List<SecretValueEntry>` | no |
| `NextToken` | `string` | no |
| `Errors` | `List<APIErrorType>` | no |

## CancelRotateSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |

## CreateSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Description` | `string` | no |
| `KmsKeyId` | `string` | no |
| `SecretBinary` | `blob` | no |
| `SecretString` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `AddReplicaRegions` | `List<ReplicaRegionType>` | no |
| `ForceOverwriteReplicaSecret` | `boolean` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |
| `ReplicationStatus` | `List<ReplicationStatusType>` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |

## DeleteSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `RecoveryWindowInDays` | `long` | no |
| `ForceDeleteWithoutRecovery` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `DeletionDate` | `timestamp` | no |

## DescribeSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `Description` | `string` | no |
| `KmsKeyId` | `string` | no |
| `RotationEnabled` | `boolean` | no |
| `RotationLambdaARN` | `string` | no |
| `RotationRules` | `RotationRulesType` | no |
| `ExternalSecretRotationMetadata` | `List<ExternalSecretRotationMetadataItem>` | no |
| `ExternalSecretRotationRoleArn` | `string` | no |
| `LastRotatedDate` | `timestamp` | no |
| `LastChangedDate` | `timestamp` | no |
| `LastAccessedDate` | `timestamp` | no |
| `DeletedDate` | `timestamp` | no |
| `NextRotationDate` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |
| `VersionIdsToStages` | `Map<List<string>>` | no |
| `OwningService` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `PrimaryRegion` | `string` | no |
| `ReplicationStatus` | `List<ReplicationStatusType>` | no |

## GetRandomPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PasswordLength` | `long` | no |
| `ExcludeCharacters` | `string` | no |
| `ExcludeNumbers` | `boolean` | no |
| `ExcludePunctuation` | `boolean` | no |
| `ExcludeUppercase` | `boolean` | no |
| `ExcludeLowercase` | `boolean` | no |
| `IncludeSpace` | `boolean` | no |
| `RequireEachIncludedType` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RandomPassword` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `ResourcePolicy` | `string` | no |

## GetSecretValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `VersionId` | `string` | no |
| `VersionStage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |
| `SecretBinary` | `blob` | no |
| `SecretString` | `string` | no |
| `VersionStages` | `List<string>` | no |
| `CreatedDate` | `timestamp` | no |

## ListSecretVersionIds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeDeprecated` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Versions` | `List<SecretVersionsListEntry>` | no |
| `NextToken` | `string` | no |
| `ARN` | `string` | no |
| `Name` | `string` | no |

## ListSecrets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IncludePlannedDeletion` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretList` | `List<SecretListEntry>` | no |
| `NextToken` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `ResourcePolicy` | `string` | yes |
| `BlockPublicPolicy` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |

## PutSecretValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `SecretBinary` | `blob` | no |
| `SecretString` | `string` | no |
| `VersionStages` | `List<string>` | no |
| `RotationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |
| `VersionStages` | `List<string>` | no |

## RemoveRegionsFromReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `RemoveReplicaRegions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `ReplicationStatus` | `List<ReplicationStatusType>` | no |

## ReplicateSecretToRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `AddReplicaRegions` | `List<ReplicaRegionType>` | yes |
| `ForceOverwriteReplicaSecret` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `ReplicationStatus` | `List<ReplicationStatusType>` | no |

## RestoreSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |

## RotateSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `RotationLambdaARN` | `string` | no |
| `RotationRules` | `RotationRulesType` | no |
| `ExternalSecretRotationMetadata` | `List<ExternalSecretRotationMetadataItem>` | no |
| `ExternalSecretRotationRoleArn` | `string` | no |
| `RotateImmediately` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |

## StopReplicationToReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Description` | `string` | no |
| `KmsKeyId` | `string` | no |
| `SecretBinary` | `blob` | no |
| `SecretString` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |
| `VersionId` | `string` | no |

## UpdateSecretVersionStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | yes |
| `VersionStage` | `string` | yes |
| `RemoveFromVersionId` | `string` | no |
| `MoveToVersionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | no |
| `Name` | `string` | no |

## ValidateResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretId` | `string` | no |
| `ResourcePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyValidationPassed` | `boolean` | no |
| `ValidationErrors` | `List<ValidationErrorsEntry>` | no |

