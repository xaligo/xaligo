# AWS Signer

API version: 2017-08-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/signer/2017-08-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddProfilePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `profileVersion` | `string` | no |
| `action` | `string` | yes |
| `principal` | `string` | yes |
| `revisionId` | `string` | no |
| `statementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revisionId` | `string` | no |

## CancelSigningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeSigningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `source` | `Source` | no |
| `signingMaterial` | `SigningMaterial` | no |
| `platformId` | `string` | no |
| `platformDisplayName` | `string` | no |
| `profileName` | `string` | no |
| `profileVersion` | `string` | no |
| `overrides` | `SigningPlatformOverrides` | no |
| `signingParameters` | `Map<string>` | no |
| `createdAt` | `timestamp` | no |
| `completedAt` | `timestamp` | no |
| `signatureExpiresAt` | `timestamp` | no |
| `requestedBy` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `revocationRecord` | `SigningJobRevocationRecord` | no |
| `signedObject` | `SignedObject` | no |
| `jobOwner` | `string` | no |
| `jobInvoker` | `string` | no |

## GetRevocationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signatureTimestamp` | `timestamp` | yes |
| `platformId` | `string` | yes |
| `profileVersionArn` | `string` | yes |
| `jobArn` | `string` | yes |
| `certificateHashes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revokedEntities` | `List<string>` | no |

## GetSigningPlatform

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `platformId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `platformId` | `string` | no |
| `displayName` | `string` | no |
| `partner` | `string` | no |
| `target` | `string` | no |
| `category` | `string` | no |
| `signingConfiguration` | `SigningConfiguration` | no |
| `signingImageFormat` | `SigningImageFormat` | no |
| `maxSizeInMB` | `integer` | no |
| `revocationSupported` | `boolean` | no |

## GetSigningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `profileOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | no |
| `profileVersion` | `string` | no |
| `profileVersionArn` | `string` | no |
| `revocationRecord` | `SigningProfileRevocationRecord` | no |
| `signingMaterial` | `SigningMaterial` | no |
| `platformId` | `string` | no |
| `platformDisplayName` | `string` | no |
| `signatureValidityPeriod` | `SignatureValidityPeriod` | no |
| `overrides` | `SigningPlatformOverrides` | no |
| `signingParameters` | `Map<string>` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `arn` | `string` | no |
| `tags` | `Map<string>` | no |

## ListProfilePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revisionId` | `string` | no |
| `policySizeBytes` | `integer` | no |
| `permissions` | `List<Permission>` | no |
| `nextToken` | `string` | no |

## ListSigningJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `platformId` | `string` | no |
| `requestedBy` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `isRevoked` | `boolean` | no |
| `signatureExpiresBefore` | `timestamp` | no |
| `signatureExpiresAfter` | `timestamp` | no |
| `jobInvoker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<SigningJob>` | no |
| `nextToken` | `string` | no |

## ListSigningPlatforms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `category` | `string` | no |
| `partner` | `string` | no |
| `target` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `platforms` | `List<SigningPlatform>` | no |
| `nextToken` | `string` | no |

## ListSigningProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeCanceled` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `platformId` | `string` | no |
| `statuses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profiles` | `List<SigningProfile>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutSigningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `signingMaterial` | `SigningMaterial` | no |
| `signatureValidityPeriod` | `SignatureValidityPeriod` | no |
| `platformId` | `string` | yes |
| `overrides` | `SigningPlatformOverrides` | no |
| `signingParameters` | `Map<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `profileVersion` | `string` | no |
| `profileVersionArn` | `string` | no |

## RemoveProfilePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `revisionId` | `string` | yes |
| `statementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revisionId` | `string` | no |

## RevokeSignature

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `jobOwner` | `string` | no |
| `reason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeSigningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `profileVersion` | `string` | yes |
| `reason` | `string` | yes |
| `effectiveTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SignPayload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileName` | `string` | yes |
| `profileOwner` | `string` | no |
| `payload` | `blob` | yes |
| `payloadFormat` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `jobOwner` | `string` | no |
| `metadata` | `Map<string>` | no |
| `signature` | `blob` | no |

## StartSigningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `source` | `Source` | yes |
| `destination` | `Destination` | yes |
| `profileName` | `string` | yes |
| `clientRequestToken` | `string` | yes |
| `profileOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `jobOwner` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


