# AWS S3 Control

API version: 2018-08-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/s3control/2018-08-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAccessGrantsIdentityCenter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `IdentityCenterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAccessGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccessGrantsLocationId` | `string` | yes |
| `AccessGrantsLocationConfiguration` | `AccessGrantsLocationConfiguration` | no |
| `Grantee` | `Grantee` | yes |
| `Permission` | `string` | yes |
| `ApplicationArn` | `string` | no |
| `S3PrefixType` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | no |
| `AccessGrantId` | `string` | no |
| `AccessGrantArn` | `string` | no |
| `Grantee` | `Grantee` | no |
| `AccessGrantsLocationId` | `string` | no |
| `AccessGrantsLocationConfiguration` | `AccessGrantsLocationConfiguration` | no |
| `Permission` | `string` | no |
| `ApplicationArn` | `string` | no |
| `GrantScope` | `string` | no |

## CreateAccessGrantsInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `IdentityCenterArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | no |
| `AccessGrantsInstanceId` | `string` | no |
| `AccessGrantsInstanceArn` | `string` | no |
| `IdentityCenterArn` | `string` | no |
| `IdentityCenterInstanceArn` | `string` | no |
| `IdentityCenterApplicationArn` | `string` | no |

## CreateAccessGrantsLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `LocationScope` | `string` | yes |
| `IAMRoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | no |
| `AccessGrantsLocationId` | `string` | no |
| `AccessGrantsLocationArn` | `string` | no |
| `LocationScope` | `string` | no |
| `IAMRoleArn` | `string` | no |

## CreateAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Bucket` | `string` | yes |
| `VpcConfiguration` | `VpcConfiguration` | no |
| `PublicAccessBlockConfiguration` | `PublicAccessBlockConfiguration` | no |
| `BucketAccountId` | `string` | no |
| `Scope` | `Scope` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointArn` | `string` | no |
| `Alias` | `string` | no |

## CreateAccessPointForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Configuration` | `ObjectLambdaConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectLambdaAccessPointArn` | `string` | no |
| `Alias` | `ObjectLambdaAccessPointAlias` | no |

## CreateBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACL` | `string` | no |
| `Bucket` | `string` | yes |
| `CreateBucketConfiguration` | `CreateBucketConfiguration` | no |
| `GrantFullControl` | `string` | no |
| `GrantRead` | `string` | no |
| `GrantReadACP` | `string` | no |
| `GrantWrite` | `string` | no |
| `GrantWriteACP` | `string` | no |
| `ObjectLockEnabledForBucket` | `boolean` | no |
| `OutpostId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Location` | `string` | no |
| `BucketArn` | `string` | no |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ConfirmationRequired` | `boolean` | no |
| `Operation` | `JobOperation` | yes |
| `Report` | `JobReport` | yes |
| `ClientRequestToken` | `string` | yes |
| `Manifest` | `JobManifest` | no |
| `Description` | `string` | no |
| `Priority` | `integer` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `List<S3Tag>` | no |
| `ManifestGenerator` | `JobManifestGenerator` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## CreateMultiRegionAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Details` | `CreateMultiRegionAccessPointInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestTokenARN` | `string` | no |

## CreateStorageLensGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `StorageLensGroup` | `StorageLensGroup` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccessGrantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessGrantsInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessGrantsInstanceResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessGrantsLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccessGrantsLocationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessPointForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessPointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessPointPolicyForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessPointScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBucketLifecycleConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBucketReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBucketTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJobTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMultiRegionAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Details` | `DeleteMultiRegionAccessPointInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestTokenARN` | `string` | no |

## DeletePublicAccessBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStorageLensConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigId` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStorageLensConfigurationTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigId` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStorageLensGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `JobDescriptor` | no |

## DescribeMultiRegionAccessPointOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RequestTokenARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AsyncOperation` | `AsyncOperation` | no |

## DissociateAccessGrantsIdentityCenter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccessGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccessGrantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | no |
| `AccessGrantId` | `string` | no |
| `AccessGrantArn` | `string` | no |
| `Grantee` | `Grantee` | no |
| `Permission` | `string` | no |
| `AccessGrantsLocationId` | `string` | no |
| `AccessGrantsLocationConfiguration` | `AccessGrantsLocationConfiguration` | no |
| `GrantScope` | `string` | no |
| `ApplicationArn` | `string` | no |

## GetAccessGrantsInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessGrantsInstanceArn` | `string` | no |
| `AccessGrantsInstanceId` | `string` | no |
| `IdentityCenterArn` | `string` | no |
| `IdentityCenterInstanceArn` | `string` | no |
| `IdentityCenterApplicationArn` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## GetAccessGrantsInstanceForPrefix

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `S3Prefix` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessGrantsInstanceArn` | `string` | no |
| `AccessGrantsInstanceId` | `string` | no |

## GetAccessGrantsInstanceResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `Organization` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## GetAccessGrantsLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccessGrantsLocationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | no |
| `AccessGrantsLocationId` | `string` | no |
| `AccessGrantsLocationArn` | `string` | no |
| `LocationScope` | `string` | no |
| `IAMRoleArn` | `string` | no |

## GetAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Bucket` | `string` | no |
| `NetworkOrigin` | `string` | no |
| `VpcConfiguration` | `VpcConfiguration` | no |
| `PublicAccessBlockConfiguration` | `PublicAccessBlockConfiguration` | no |
| `CreationDate` | `timestamp` | no |
| `Alias` | `string` | no |
| `AccessPointArn` | `string` | no |
| `Endpoints` | `Map<string>` | no |
| `BucketAccountId` | `string` | no |
| `DataSourceId` | `string` | no |
| `DataSourceType` | `string` | no |

## GetAccessPointConfigurationForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `ObjectLambdaConfiguration` | no |

## GetAccessPointForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `PublicAccessBlockConfiguration` | `PublicAccessBlockConfiguration` | no |
| `CreationDate` | `timestamp` | no |
| `Alias` | `ObjectLambdaAccessPointAlias` | no |

## GetAccessPointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetAccessPointPolicyForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetAccessPointPolicyStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyStatus` | `PolicyStatus` | no |

## GetAccessPointPolicyStatusForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyStatus` | `PolicyStatus` | no |

## GetAccessPointScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `Scope` | no |

## GetBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bucket` | `string` | no |
| `PublicAccessBlockEnabled` | `boolean` | no |
| `CreationDate` | `timestamp` | no |

## GetBucketLifecycleConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<LifecycleRule>` | no |

## GetBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetBucketReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfiguration` | `ReplicationConfiguration` | no |

## GetBucketTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSet` | `List<S3Tag>` | yes |

## GetBucketVersioning

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `MFADelete` | `string` | no |

## GetDataAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Target` | `string` | yes |
| `Permission` | `string` | yes |
| `DurationSeconds` | `integer` | no |
| `Privilege` | `string` | no |
| `TargetType` | `string` | no |
| `AuditContext` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `MatchedGrantTarget` | `string` | no |
| `Grantee` | `Grantee` | no |

## GetJobTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<S3Tag>` | no |

## GetMultiRegionAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPoint` | `MultiRegionAccessPointReport` | no |

## GetMultiRegionAccessPointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `MultiRegionAccessPointPolicyDocument` | no |

## GetMultiRegionAccessPointPolicyStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Established` | `PolicyStatus` | no |

## GetMultiRegionAccessPointRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Mrap` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mrap` | `string` | no |
| `Routes` | `List<MultiRegionAccessPointRoute>` | no |

## GetPublicAccessBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicAccessBlockConfiguration` | `PublicAccessBlockConfiguration` | no |

## GetStorageLensConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigId` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageLensConfiguration` | `StorageLensConfiguration` | no |

## GetStorageLensConfigurationTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigId` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<StorageLensTag>` | no |

## GetStorageLensGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageLensGroup` | `StorageLensGroup` | no |

## ListAccessGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `GranteeType` | `string` | no |
| `GranteeIdentifier` | `string` | no |
| `Permission` | `string` | no |
| `GrantScope` | `string` | no |
| `ApplicationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AccessGrantsList` | `List<ListAccessGrantEntry>` | no |

## ListAccessGrantsInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AccessGrantsInstancesList` | `List<ListAccessGrantsInstanceEntry>` | no |

## ListAccessGrantsLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LocationScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AccessGrantsLocationsList` | `List<ListAccessGrantsLocationsEntry>` | no |

## ListAccessPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DataSourceId` | `string` | no |
| `DataSourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointList` | `List<AccessPoint>` | no |
| `NextToken` | `string` | no |

## ListAccessPointsForDirectoryBuckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `DirectoryBucket` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointList` | `List<AccessPoint>` | no |
| `NextToken` | `string` | no |

## ListAccessPointsForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectLambdaAccessPointList` | `List<ObjectLambdaAccessPoint>` | no |
| `NextToken` | `string` | no |

## ListCallerAccessGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `GrantScope` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AllowedByApplication` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CallerAccessGrantsList` | `List<ListCallerAccessGrantsEntry>` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobStatuses` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Jobs` | `List<JobListDescriptor>` | no |

## ListMultiRegionAccessPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPoints` | `List<MultiRegionAccessPointReport>` | no |
| `NextToken` | `string` | no |

## ListRegionalBuckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `OutpostId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionalBucketList` | `List<RegionalBucket>` | no |
| `NextToken` | `string` | no |

## ListStorageLensConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StorageLensConfigurationList` | `List<ListStorageLensConfigurationEntry>` | no |

## ListStorageLensGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StorageLensGroupList` | `List<ListStorageLensGroupEntry>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutAccessGrantsInstanceResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Policy` | `string` | yes |
| `Organization` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `Organization` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## PutAccessPointConfigurationForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Configuration` | `ObjectLambdaConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccessPointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccessPointPolicyForObjectLambda

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccessPointScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `Scope` | `Scope` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBucketLifecycleConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |
| `LifecycleConfiguration` | `LifecycleConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |
| `ConfirmRemoveSelfBucketAccess` | `boolean` | no |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBucketReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |
| `ReplicationConfiguration` | `ReplicationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBucketTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |
| `Tagging` | `Tagging` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBucketVersioning

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Bucket` | `string` | yes |
| `MFA` | `string` | no |
| `VersioningConfiguration` | `VersioningConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutJobTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobId` | `string` | yes |
| `Tags` | `List<S3Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMultiRegionAccessPointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Details` | `PutMultiRegionAccessPointPolicyInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestTokenARN` | `string` | no |

## PutPublicAccessBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicAccessBlockConfiguration` | `PublicAccessBlockConfiguration` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutStorageLensConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigId` | `string` | yes |
| `AccountId` | `string` | yes |
| `StorageLensConfiguration` | `StorageLensConfiguration` | yes |
| `Tags` | `List<StorageLensTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutStorageLensConfigurationTagging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigId` | `string` | yes |
| `AccountId` | `string` | yes |
| `Tags` | `List<StorageLensTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubmitMultiRegionAccessPointRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Mrap` | `string` | yes |
| `RouteUpdates` | `List<MultiRegionAccessPointRoute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccessGrantsLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccessGrantsLocationId` | `string` | yes |
| `IAMRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | no |
| `AccessGrantsLocationId` | `string` | no |
| `AccessGrantsLocationArn` | `string` | no |
| `LocationScope` | `string` | no |
| `IAMRoleArn` | `string` | no |

## UpdateJobPriority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobId` | `string` | yes |
| `Priority` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `Priority` | `integer` | yes |

## UpdateJobStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `JobId` | `string` | yes |
| `RequestedJobStatus` | `string` | yes |
| `StatusUpdateReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `Status` | `string` | no |
| `StatusUpdateReason` | `string` | no |

## UpdateStorageLensGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AccountId` | `string` | yes |
| `StorageLensGroup` | `StorageLensGroup` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


